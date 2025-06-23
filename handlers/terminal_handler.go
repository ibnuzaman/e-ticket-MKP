package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"e-ticketing/models"
)

type TerminalHandler struct {
	DB *gorm.DB
}

func NewTerminalHandler(db *gorm.DB) *TerminalHandler {
	return &TerminalHandler{DB: db}
}

type CreateTerminalRequest struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (h *TerminalHandler) CreateTerminal(c echo.Context) error {
	var req CreateTerminalRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	userID, ok := c.Get("user_id").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
	}
	role, ok := c.Get("role").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid role"})
	}

	if role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Admin access required"})
	}

	terminal := models.Terminal{
		TerminalID: uuid.New(),
		Name:       req.Name,
		Location:   req.Location,
		Status:     "active",
	}

	var existing models.Terminal
	if err := h.DB.Where("name = ?", terminal.Name).First(&existing).Error; err == nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Terminal with this name already exists"})
	} else if err != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	if err := h.DB.Create(&terminal).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create terminal"})
	}

	log.Printf("Terminal created by user %s: %s", userID, terminal.Name)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"terminal_id": terminal.TerminalID,
		"name":        terminal.Name,
		"location":    terminal.Location,
		"status":      terminal.Status,
	})
}
