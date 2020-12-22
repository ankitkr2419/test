CREATE TYPE cartridge_type AS ENUM('extraction', 'pcr');

CREATE TABLE IF NOT EXISTS cartridges(
 labware_id int NOT NULL,
 type cartridge_type DEFAULT 'extraction',
 description varchar(250),
 well_num int NOT NULL,
 distance decimal NOT NULL,
 height decimal NOT NULL,
 volume decimal NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (labware_id) REFERENCES labwares(id) ON UPDATE CASCADE ON DELETE CASCADE,
 PRIMARY KEY (labware_id, type, well_num));