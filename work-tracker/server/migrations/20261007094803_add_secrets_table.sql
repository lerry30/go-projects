-- +goose Up
CREATE TABLE secrets(
    employee_id INTEGER REFERENCES users(id),
    secret_key VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY
);

-- +goose Down
DROP TABLE secrets;
