CREATE TABLE IF NOT EXISTS heating(
 id uuid primary key,
 temperature int,
 follow_temp boolean DEFAULT FALSE,
 duration int,
 process_id uuid,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (process_id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
