CREATE TYPE tip_movement_type AS ENUM ('cartridge_1', 'deck', 'cartridge_2');

CREATE TABLE  tip_movement(
 id uuid PRIMARY KEY,
 type tip_movement_type DEFAULT 'deck',
 position int CHECK (position<10),
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
