CREATE TYPE tip_tube_type AS ENUM('tip', 'tube');

CREATE TABLE IF NOT EXISTS tips_and_tubes(
id int PRIMARY KEY,
name varchar(50) UNIQUE NOT NULL,
type tip_tube_type,
allowed_positions int[],
volume decimal NOT NULL,
height decimal NOT NULL,
created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
