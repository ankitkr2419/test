CREATE TYPE cartridge_type AS ENUM('extraction', 'pcr');

CREATE TABLE IF NOT EXISTS cartridges(
 labware_id int NOT NULL,
 type cartridge_type DEFAULT 'extraction',
 description varchar(250) NOT NULL,
 wells int DEFAULT 8 NOT NULL,
 distances decimal[] NOT NULL,
 heights decimal[] NOT NULL,
 volumes decimal[] NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (labware_id) REFERENCES labwares(id),
 PRIMARY KEY (labware_id, type));