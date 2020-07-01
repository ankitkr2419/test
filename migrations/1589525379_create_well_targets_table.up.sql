CREATE TABLE well_targets (
  well_id uuid,
  target_id uuid,
  ct varchar(50) DEFAULT '',
  FOREIGN KEY (target_id) REFERENCES targets(id),
  FOREIGN KEY (well_id) REFERENCES wells(id) ON DELETE CASCADE
);
