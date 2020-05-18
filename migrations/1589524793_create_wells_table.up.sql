CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE wells (
  id uuid primary key default uuid_generate_v4(),
  position integer,
  experiment_id uuid,
  sample_id uuid,
  task varchar(50),
  color_code float
);
