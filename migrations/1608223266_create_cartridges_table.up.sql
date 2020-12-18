CREATE TYPE cartridge_type AS ENUM('extraction', 'pcr');

CREATE TABLE IF NOT EXISTS cartridges(
 number serial primary key,
 type cartridge_type DEFAULT 'extraction',
 wells int DEFAULT 8 NOT NULL,
 distances decimal[] NOT NULL,
 heights decimal[] NOT NULL,
 volumes decimal[] NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
CONSTRAINT distances_array_len CHECK(wells = array_length(distances, 1) ),
CONSTRAINT heights_array_len CHECK( wells = array_length(heights, 1) ),
CONSTRAINT volumes_array_len CHECK( wells = array_length(volumes, 1) )
);
