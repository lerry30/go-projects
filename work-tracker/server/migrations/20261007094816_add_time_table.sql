-- +goose Up
CREATE TABLE time(
    employee_id INTEGER REFERENCES users(id),
    tracker_id INTEGER REFERENCES tracker(id),
    login TIME NOT NULL,
    first_break_out TIME NOT NULL,
    first_break_in TIME NOT NULL,
    lunch_out TIME NOT NULL,
    lunch_in TIME NOT NULL,
    second_break_out TIME NOT NULL,
    second_break_in TIME NOT NULL,
    logout TIME NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY
);

-- +goose Down
DROP TABLE time;
