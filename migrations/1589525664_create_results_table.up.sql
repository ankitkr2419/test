CREATE TABLE results (
  experiment_id uuid,
  target_id uuid,
  well_position integer,
  cycle integer,
  f_value float,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (experiment_id) REFERENCES experiments(id) ON DELETE CASCADE
);
