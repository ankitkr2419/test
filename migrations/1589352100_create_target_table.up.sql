CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE targets (
  id uuid primary key default uuid_generate_v4(),
	name varchar(50),
  dye_id uuid
);
