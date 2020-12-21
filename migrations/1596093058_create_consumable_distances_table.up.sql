CREATE TABLE IF NOT EXISTS consumable_distances(
name varchar(50) PRIMARY KEY,
description varchar(250),
deck_1_distance decimal NULL,
created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP );
