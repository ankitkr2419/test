CREATE TYPE discard_type AS ENUM ('at_pickup_passing', 'at_discard_box');

CREATE TABLE IF NOT EXISTS piercing(
 id uuid,
 cartridge_ids int[] NOT NULL,
 discard discard_type DEFAULT 'at_pickup_passing',
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (id),
 FOREIGN KEY (id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);