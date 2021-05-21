ALTER TYPE process_type ADD VALUE IF NOT EXISTS 'TipPickup';
ALTER TYPE process_type RENAME VALUE 'TipOperation' TO 'TipDiscard';