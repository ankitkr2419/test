ALTER TABLE IF EXISTS piercing ADD COLUMN discard discard_type DEFAULT 'at_pickup_passing';
ALTER TABLE IF EXISTS tip_operation DROP COLUMN discard;
