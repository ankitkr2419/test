CREATE TABLE template_targets
(
    template_id uuid,
    target_id uuid,
    threshold float,
    CONSTRAINT unqexp UNIQUE (template_id, target_id));