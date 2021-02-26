CREATE TYPE magnet_movement_type AS ENUM ('attach', 'detach');

CREATE TABLE  magnet(
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 type magnet_movement_type DEFAULT 'attach',
 process_id uuid,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (process_id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
