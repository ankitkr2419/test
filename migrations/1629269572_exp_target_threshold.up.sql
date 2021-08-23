CREATE TABLE IF NOT EXISTS exp_target_threshold(
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 exp_id uuid,
 target_id uuid,
 threshold DECIMAL,
 CONSTRAINT unett UNIQUE (exp_id,target_id),
 FOREIGN KEY (exp_id) REFERENCES experiments(id) ON DELETE CASCADE ON UPDATE CASCADE,
 FOREIGN KEY (target_id) REFERENCES targets(id) ON DELETE CASCADE ON UPDATE CASCADE);
