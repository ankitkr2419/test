CREATE TYPE labware_type AS ENUM('tip', 'tube');
CREATE TYPE labware_subtype AS ENUM('extraction', 'pcr', 'piercing', 'sample' , 'shaker');

CREATE TABLE IF NOT EXISTS labwares(
id int NOT NULL,
name varchar(50) NOT NULL, 
type labware_type,
subtype labware_subtype,
volume decimal NOT NULL,
height decimal NOT NULL,
created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (id, name),
FOREIGN KEY (name) REFERENCES consumable_distances(name));
