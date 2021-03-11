CREATE TABLE IF NOT EXISTS shaking(
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 with_temp boolean DEFAULT FALSE,
 follow_temp boolean DEFAULT FALSE,
 temperature int,
 rpm_1 int,
 rpm_2 int,
 time_1 int,
 time_2 int,
 process_id uuid UNIQUE NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
FOREIGN KEY (process_id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
