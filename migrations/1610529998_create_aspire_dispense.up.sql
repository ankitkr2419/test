CREATE TYPE aspire_dispense_type AS ENUM ('well_to_shaker', 'shaker_to_well', 'well_to_well', 'well_to_deck', 'deck_to_well', 'deck_to_deck');
 
CREATE TABLE IF NOT EXISTS aspire_dispense(
 id uuid PRIMARY KEY,
 category aspire_dispense_type,
 well_no_source int,
 aspire_height float,
 aspire_mixing_volume float,
 aspire_no_of_cycles int,
 aspire_volume float,
 aspire_air_volume float,
 dispense_height float,
 dispense_mixing_volume float,
 dispense_no_of_cycles int,
 dispense_vol float,
 dispense_blow float,
 well_no_destination int,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
 
