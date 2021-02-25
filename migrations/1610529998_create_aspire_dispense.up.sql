CREATE TYPE aspire_dispense_type AS ENUM ('well_to_shaker', 'shaker_to_well', 'well_to_well', 'well_to_deck', 'deck_to_well', 'deck_to_deck');
CREATE TYPE aspire_dispense_cartridge_type AS ENUM('cartridge_1', 'cartridge_2');

CREATE TABLE IF NOT EXISTS aspire_dispense(
 id uuid PRIMARY KEY,
 category aspire_dispense_type,
 cartridge_type aspire_dispense_cartridge_type,
 source_position int,
 aspire_height float,
 aspire_mixing_volume float,
 aspire_no_of_cycles int,
 aspire_volume float,
 aspire_air_volume float,
 dispense_height float,
 dispense_mixing_volume float,
 dispense_no_of_cycles int,
 dispense_volume float,
 dispense_blow_volume float,
 destination_position int,
 process_id uuid,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (process_id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
 
