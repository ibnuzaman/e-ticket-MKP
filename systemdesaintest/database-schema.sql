CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'operator')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cards (
    card_id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    balance DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    status VARCHAR(20) NOT NULL CHECK (status IN ('active', 'inactive', 'blocked')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE terminals (
    terminal_id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location VARCHAR(255) NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('active', 'inactive')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tariffs (
    id UUID PRIMARY KEY,
    origin_terminal_id UUID REFERENCES terminals(terminal_id),
    destination_terminal_id UUID REFERENCES terminals(terminal_id),
    amount DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    transaction_id UUID PRIMARY KEY,
    card_id UUID REFERENCES cards(card_id),
    check_in_terminal_id UUID REFERENCES terminals(terminal_id),
    check_out_terminal_id UUID REFERENCES terminals(terminal_id),
    check_in_time TIMESTAMP,
    check_out_time TIMESTAMP,
    amount DECIMAL(10,2),
    status VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'completed', 'failed')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_cards_card_id ON cards(card_id);
CREATE INDEX idx_transactions_card_id_check_in_time ON transactions(card_id, check_in_time);
CREATE INDEX idx_tariffs_origin_destination ON tariffs(origin_terminal_id, destination_terminal_id);