CREATE TYPE tip_dock_type AS ENUM ('cartridge_1', 'deck', 'cartridge_2');

CREATE TABLE  tip_dock(
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 type tip_dock_type DEFAULT 'deck',
 position int CHECK (position<10),
 height decimal,
 process_id uuid,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (process_id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
