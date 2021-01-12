CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE IF NOT EXISTS recipes(
 id uuid primary key default uuid_generate_v4(),
 name varchar(50) NOT NULL,
 description varchar(200) NOT NULL,
 labware_id int NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (labware_id) references labwares(id) ON UPDATE CASCADE ON DELETE CASCADE);