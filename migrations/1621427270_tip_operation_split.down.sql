-- NOTE: No provision provided by postgresql to DROP a single value
-- This not dropping of 'TipPickup' won't cause any issue
ALTER TYPE process_type RENAME VALUE 'TipDiscard' TO 'TipOperation';
UPDATE processes SET type='TipOperation' where type='TipPickup';