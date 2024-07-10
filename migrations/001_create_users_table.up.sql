CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    passport_serie VARCHAR(4) NOT NULL,
    passport_number VARCHAR(6) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    patronymic VARCHAR(50),
    address TEXT
);

CREATE UNIQUE INDEX idx_users_passport ON users (passport_serie, passport_number);