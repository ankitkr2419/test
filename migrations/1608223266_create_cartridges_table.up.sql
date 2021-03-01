CREATE TYPE cartridge_type AS ENUM('cartridge_1', 'cartridge_2');
 
CREATE TABLE IF NOT EXISTS cartridges(
 id int PRIMARY KEY,
 type cartridge_type DEFAULT 'cartridge_1',
 description varchar(250),
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
