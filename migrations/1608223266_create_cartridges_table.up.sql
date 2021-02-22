CREATE TYPE cartridge_type AS ENUM('cartridge_1', 'cartridge_2');
 
CREATE TABLE IF NOT EXISTS cartridges(
 id int NOT NULL,
 type cartridge_type DEFAULT 'cartridge_1',
 description varchar(250),
 well_num int NOT NULL,
 distance decimal NOT NULL,
 height decimal NOT NULL,
 volume decimal NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (id, well_num)
);
