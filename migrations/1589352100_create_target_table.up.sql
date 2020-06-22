CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE targets (
  id uuid primary key default uuid_generate_v4(),
	name varchar(50),
  dye_id uuid,
  CONSTRAINT unqtarget UNIQUE(name, dye_id),
  FOREIGN KEY (dye_id) REFERENCES dyes(id));
