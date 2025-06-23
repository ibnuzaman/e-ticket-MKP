package seeders

import (
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"e-ticketing/models"
)

func SeedUsers(db *gorm.DB) error {
	users := []struct {
		ID       uuid.UUID
		Username string
		Password string
		Role     string
	}{
		{
			ID:       uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
			Username: "admin",
			Password: "password",
			Role:     "admin",
		},
	}

	for _, user := range users {
		var existingUser models.User
		if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
			log.Printf("User %s already exists, skipping seeding", user.Username)
			continue
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		newUser := models.User{
			ID:           user.ID,
			Username:     user.Username,
			PasswordHash: string(hashedPassword),
			Role:         user.Role,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		if err := db.Create(&newUser).Error; err != nil {
			return err
		}
		log.Printf("Seeded user: %s", user.Username)
	}

	return nil
}
