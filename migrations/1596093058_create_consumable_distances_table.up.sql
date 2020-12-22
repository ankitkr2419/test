CREATE TABLE IF NOT EXISTS consumable_distances(
 id int PRIMARY KEY,
 name varchar(50) UNIQUE NOT NULL,
 description varchar(250),
 deck_1_distance decimal NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP );
