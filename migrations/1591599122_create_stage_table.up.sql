CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE stages
(
    id uuid primary key default uuid_generate_v4(),
    name varchar(50),
    type varchar(50),
    repeat_count integer,
    template_id uuid,
    FOREIGN KEY (template_id) REFERENCES templates(id) ON DELETE CASCADE
);
