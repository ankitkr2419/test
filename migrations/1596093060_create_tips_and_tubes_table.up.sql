CREATE TABLE IF NOT EXISTS tips_and_tubes(
labware_id int NOT NULL,
name varchar(50) NOT NULL,
volume decimal NOT NULL,
height decimal NOT NULL,
created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (labware_id, name),
FOREIGN KEY (name) REFERENCES consumable_distances(name),
FOREIGN KEY (labware_id) REFERENCES labwares(id));
