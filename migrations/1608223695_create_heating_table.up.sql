CREATE TABLE IF NOT EXISTS heating(
 id uuid,
 temperature decimal,
 follow_temp boolean DEFAULT FALSE,
 duration decimal,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);