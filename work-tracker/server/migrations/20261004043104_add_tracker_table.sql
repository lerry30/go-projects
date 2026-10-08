-- +goose Up
CREATE TABLE tracker(
    employee_id INTEGER REFERENCES users(id),
    spread_sheet_id VARCHAR(255) UNIQUE NOT NULL,
    year SMALLINT NOT NULL,
    month SMALLINT NOT NULL CHECK (month BETWEEN 1 AND 12),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY
);

-- +goose Down
DROP TABLE tracker;