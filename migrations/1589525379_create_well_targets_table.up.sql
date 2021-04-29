CREATE TABLE IF NOT EXISTS well_targets (
  experiment_id uuid,
  well_position integer,
  target_id uuid,
  selected boolean DEFAULT false,
  ct varchar(50) DEFAULT '',
  FOREIGN KEY (target_id) REFERENCES targets(id),
  FOREIGN KEY (experiment_id) REFERENCES experiments(id),
  CONSTRAINT unqexptargets UNIQUE (well_position, target_id,experiment_id));
