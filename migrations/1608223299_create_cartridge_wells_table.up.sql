CREATE TABLE IF NOT EXISTS cartridge_wells(
 id int,
 well_num int NOT NULL,
 distance decimal NOT NULL,
 height decimal NOT NULL,
 volume decimal NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (id, well_num),
 FOREIGN KEY (id) REFERENCES cartridges(id)
);
