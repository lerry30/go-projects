-- +goose Up
CREATE TABLE appscripts(
    employee_id INTEGER REFERENCES users(id),
    appscript_endpoint VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY
);

-- +goose Down
DROP TABLE appscripts;
