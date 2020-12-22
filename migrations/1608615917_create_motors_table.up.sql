CREATE TABLE IF NOT EXISTS motors(
 number int PRIMARY KEY,
 name varchar(50) NOT NULL,
 ramp int default 100 NOT NULL,
 steps int default 200 NOT NULL,
 slow int default 500 NOT NULL,
 fast int default 2000 NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP);
