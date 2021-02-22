CREATE TABLE IF NOT EXISTS heating(
 id uuid primary key,
 temperature decimal,
 follow_temp boolean DEFAULT FALSE,
 duration decimal,
 shaker int,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);