-- +goose Up
CREATE TABLE appscripts(
    employee_id INTEGER REFERENCES users(id) UNIQUE,
    appscript_endpoint VARCHAR(255) NOT NULL,
    status VARCHAR(10) NOT NULL DEFAULT 'stopped' CHECK (status IN ('stopped', 'running')),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY
);

-- +goose Down
DROP TABLE appscripts;
