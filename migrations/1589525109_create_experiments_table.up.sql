CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE experiments (
  id uuid primary key default uuid_generate_v4(),
  description varchar(50),
  template_id uuid,
  operator_name varchar(50),
  start_time timestamp,
  end_time timestamp
);
