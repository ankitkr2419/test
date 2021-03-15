CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TYPE process_type AS 
ENUM('Piercing','TipOperation','TipDocking','AspireDispense','Heating', 'Shaking', 'AttachDetach', 'Delay');

CREATE TABLE IF NOT EXISTS processes(
 id uuid PRIMARY KEY default uuid_generate_v4(),
 type process_type,
 recipe_id uuid ,
 sequence_num int,
 name varchar(50) NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 CONSTRAINT unqsequence UNIQUE(recipe_id, sequence_num),
 FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON UPDATE CASCADE ON DELETE CASCADE);