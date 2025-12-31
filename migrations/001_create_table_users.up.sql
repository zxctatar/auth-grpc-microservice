CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    middle_name VARCHAR(255),
    last_name VARCHAR(255) NOT NULL,
    hash_password TEXT NOT NULL,
    email VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_email ON users(email);