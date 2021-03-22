CREATE TABLE IF NOT EXISTS heating(
 id uuid primary key DEFAULT uuid_generate_v4(),
 temperature decimal,
 follow_temp boolean DEFAULT FALSE,
 duration int,
 process_id uuid UNIQUE NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (process_id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
