CREATE TABLE IF NOT EXISTS experiment_template_targets
(
    experiment_id uuid,
    template_id uuid,
    target_id uuid,
    threshold float,
    CONSTRAINT unqexp UNIQUE (template_id, target_id,experiment_id));