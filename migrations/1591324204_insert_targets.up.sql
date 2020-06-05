INSERT INTO targets
    (
    name,dye_id)
VALUES
    ('A', '456c3482-53ef-4ccb-bf3d-41d67b7354dc'),
    ('B', '456c3482-53ef-4ccb-bf3d-41d67b7354dc'),
    ('C', '456c3482-53ef-4ccb-bf3d-41d67b7354dc'),
    ('D', '456c3482-53ef-4ccb-bf3d-41d67b7354dc'),
    ('E', '456c3482-53ef-4ccb-bf3d-41d67b7354dc'),
    ('E', '456c3482-53ef-4ccb-bf3d-41d67b7354dc'),
    ('F', '456c3482-53ef-4ccb-bf3d-41d67b7354dc'),
    ('G', '456c3482-53ef-4ccb-bf3d-41d67b7354dc'),
    ('H', '456c3482-53ef-4ccb-bf3d-41d67b7354dc')
ON CONFLICT DO NOTHING;
