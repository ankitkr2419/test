CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE steps
(
    id uuid primary key default uuid_generate_v4(),
    stage_id uuid,
    ramp_rate float,
    target_temp float,
    hold_time integer,
    data_capture boolean,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (stage_id) REFERENCES stages(id) ON DELETE CASCADE
);
