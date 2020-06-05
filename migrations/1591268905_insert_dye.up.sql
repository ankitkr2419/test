INSERT INTO dyes
    (
    name,position)
VALUES
    ('VIC',1),
    ('TAMRA',2),
    ('SYBR',3),
    ('ROX',4),
    ('OTHER',5),
    ('NFQ-MGB',6),
    ('NED',6),
    ('MUSTANG PURPLE',5),
    ('JUN',4),
    ('FAM',3),
    ('CY5',4),
    ('CM_DYE',2),
    ('ABY',1)
ON CONFLICT DO NOTHING;
