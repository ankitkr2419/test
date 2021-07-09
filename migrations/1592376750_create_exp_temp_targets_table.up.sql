CREATE TABLE IF NOT EXISTS experiment_template_targets
(
    experiment_id uuid,
    template_id uuid,
    target_id uuid,
    threshold float,
    FOREIGN KEY (experiment_id) REFERENCES experiments(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (template_id) REFERENCES templates(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (target_id) REFERENCES targets(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT unqexp UNIQUE (template_id, target_id,experiment_id));