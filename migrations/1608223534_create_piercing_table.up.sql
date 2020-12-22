CREATE TYPE discard_type AS ENUM ('at_pickup_passing', 'at_discard_box');

CREATE TABLE IF NOT EXISTS piercing(
 id uuid,
 cartridge_id int,
 type cartridge_type,
 well_num integer NOT NULL,
 discard discard_type DEFAULT 'at_pickup_passing',
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (id, well_num),
 FOREIGN KEY (id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE,
 FOREIGN KEY (cartridge_id, type) REFERENCES cartridges(labware_id, type) ON UPDATE CASCADE ON DELETE CASCADE);