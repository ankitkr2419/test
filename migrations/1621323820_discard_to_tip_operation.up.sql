ALTER TABLE IF EXISTS piercing DROP COLUMN discard;
ALTER TABLE IF EXISTS tip_operation ADD COLUMN discard discard_type DEFAULT 'at_pickup_passing';
