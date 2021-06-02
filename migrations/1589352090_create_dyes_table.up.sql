CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE IF NOT EXISTS dyes (
  id uuid primary key default uuid_generate_v4(),
  name varchar(50),
  position int,
  CONSTRAINT unqdye UNIQUE(name, position));
