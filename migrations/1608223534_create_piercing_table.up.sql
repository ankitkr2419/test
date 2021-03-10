CREATE TYPE discard_type AS ENUM ('at_pickup_passing', 'at_discard_box');

CREATE TABLE IF NOT EXISTS piercing(
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 type cartridge_type DEFAULT 'cartridge_1',
 cartridge_wells int[] NOT NULL,
 discard discard_type DEFAULT 'at_pickup_passing',
 process_id uuid UNIQUE NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (process_id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
