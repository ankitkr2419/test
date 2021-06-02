CREATE TYPE aspire_dispense_type AS ENUM ('well_to_shaker', 'shaker_to_well', 'well_to_well', 'well_to_deck', 'deck_to_well', 'deck_to_deck', 'shaker_to_deck', 'deck_to_shaker');
CREATE TYPE aspire_dispense_cartridge_type AS ENUM('cartridge_1', 'cartridge_2');

CREATE TABLE IF NOT EXISTS aspire_dispense(
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 category aspire_dispense_type,
 cartridge_type aspire_dispense_cartridge_type DEFAULT 'cartridge_1',
 source_position int,
 aspire_height float,
 aspire_mixing_volume float,
 aspire_no_of_cycles int,
 aspire_volume float,
 aspire_air_volume float,
 dispense_height float,
 dispense_mixing_volume float,
 dispense_no_of_cycles int,
 destination_position int,
 process_id uuid UNIQUE NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (process_id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
 