CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
 
CREATE TABLE IF NOT EXISTS recipes(
 id uuid primary key default uuid_generate_v4(),
 name varchar(50) NOT NULL,
 description varchar(200) NOT NULL,
 pos_1 int NOT NULL,
 pos_2 int NOT NULL,
 pos_3 int NOT NULL,
 pos_4 int NOT NULL,
 pos_5 int NOT NULL,
 pos_cartridge_1 int NOT NULL,
 pos_7 int NOT NULL,
 pos_cartridge_2 int NOT NULL,
 pos_9 int NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY(pos_1) REFERENCES tips_and_tubes(id),
 FOREIGN KEY(pos_2) REFERENCES tips_and_tubes(id),
 FOREIGN KEY(pos_3) REFERENCES tips_and_tubes(id),
 FOREIGN KEY(pos_4) REFERENCES tips_and_tubes(id),
 FOREIGN KEY(pos_5) REFERENCES tips_and_tubes(id),
 FOREIGN KEY(pos_cartridge_1) REFERENCES cartridges(id),
 FOREIGN KEY(pos_7) REFERENCES tips_and_tubes(id),
 FOREIGN KEY(pos_cartridge_2) REFERENCES cartridges(id), 
 FOREIGN KEY(pos_9) REFERENCES tips_and_tubes(id)
);
