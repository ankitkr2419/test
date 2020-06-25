CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE stages
(
    id uuid primary key default uuid_generate_v4(),
    type varchar(50),
    repeat_count integer,
    template_id uuid,
    step_count integer  NOT NULL,
    FOREIGN KEY (template_id) REFERENCES templates(id) ON DELETE CASCADE
);
