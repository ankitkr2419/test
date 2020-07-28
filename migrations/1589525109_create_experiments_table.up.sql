CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE experiments (
  id uuid primary key default uuid_generate_v4(),
  description varchar(50),
  template_id uuid,
  operator_name varchar(50),
  start_time timestamp,
  end_time timestamp,
  result varchar(50),
  repeat_cycle integer,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (template_id) REFERENCES templates(id)
);
