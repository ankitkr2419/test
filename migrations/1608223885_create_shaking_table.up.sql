CREATE TABLE IF NOT EXISTS shaking(
 id uuid PRIMARY KEY,
 set_temperature boolean DEFAULT FALSE,
 follow_temp boolean DEFAULT FALSE,
 temperature int,
 rpm_1 int,
 rpm_2 int,
 time_1 int,
 time_2 int,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
FOREIGN KEY (id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
