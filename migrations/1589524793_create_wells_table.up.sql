CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE IF NOT EXISTS wells (
  id uuid primary key default uuid_generate_v4(),
  position integer,
  experiment_id uuid,
  sample_id uuid,
  task varchar(50),
  color_code varchar(50) DEFAULT '',
  CONSTRAINT unqwell UNIQUE (experiment_id, position)
);
