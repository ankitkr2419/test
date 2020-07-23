CREATE TABLE experiment_temperatures
(
    experiment_id uuid,
    temp float,
    lid_temp float,
    cycle integer,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (experiment_id) REFERENCES experiments(id) ON DELETE CASCADE
);