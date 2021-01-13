CREATE TYPE category_type AS ENUM ('well_to_shaker', 'shaker_to_well', 'well_to_well');

CREATE TABLE IF NOT EXISTS aspire_dispense(
 id uuid,
 category category_type,
 well_no_source int,
 aspire_height float,
 aspire_mixing_volume float,
 aspire_no_of_cycles int,
 aspire_volume float,
 dispense_height float,
 dispense_mixing_volume float,
 dispense_no_of_cycles int,
 dispense_vol float,
 dispense_blow float,
 well_to_destination int,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (id),
 FOREIGN KEY (id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
