-- Insert into recipe
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2de4','covid', 'Recipe for covid extraction', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2de5','tip_docking', 'Recipe for tip_docking position on deck', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2de6','tip_docking', 'Recipe for tip_docking position on cartridge 1', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2de7','tip_docking', 'Recipe for tip_docking position on cartridge 2', 1, 2, 3, 4, 5, 1, 6, 2, 7);


--  Insert into processes
--  For Cartridge_1
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a5d058e3-7ce3-4a42-b2da-690e47139612','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 1, 'AD-WW-c1-1-2' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('4fa5c4e3-699c-47bb-ac7a-b26d04efaeb5','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 2, 'AD-WS-c1-1' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('fb88bada-ced7-4fa2-b845-4bb91e74341e','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 3, 'AD-SW-c1-2' );        
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('38a20548-85bf-4c3b-82bc-4d87b87f2dbe','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 4, 'AD-DW-c1-4-2' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('3b47f63e-fb02-460b-864f-edc4a302af5a','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 5, 'AD-WD-c1-2-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('8207aa73-74b1-4bca-86ae-88e843ef1eac','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 6, 'AD-DD-4-5' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd767','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 7, 'AD-SD-7' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd768','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 8, 'AD-DS-7' );

INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd769','TipDocking', '6b7fcfa2-8337-4d79-829a-e9bd486a2de5', 9, 'TD-DECK-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd779','TipDocking', '6b7fcfa2-8337-4d79-829a-e9bd486a2de6', 10, 'TD-C1-4' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd789','TipDocking', '6b7fcfa2-8337-4d79-829a-e9bd486a2de7', 11, 'TD-C1-5' );



--  For cartridge_2
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('bbe172f0-e315-4ebd-83ab-c1c0b531e2f8','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 9, 'AD-WW-c2-1-2' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('cd446adb-f7ed-4368-aa65-b0e04fdd2c81','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 10, 'AD-WS-c2-1' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('e6e7f15e-a06d-4d04-ac60-4de113fd8c42','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 11, 'AD-SW-c2-2' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('62c1afc8-87cb-4bb6-833e-b5716cbf71cd','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 12, 'AD-DW-c2-4-2' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3b7da1c-a402-4b81-9e76-eeac4782b779','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 13, 'AD-WD-c2-2-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd766','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 14, 'AD-DD-4-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd769','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 15, 'AD-SD-7' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd770','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 16, 'AD-DS-7' );

-- Insert into aspire_dispense process
--  For cartridge_1
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('a5d058e3-7ce3-4a42-b2da-690e47139612','a5d058e3-7ce3-4a42-b2da-690e47139612', 'well_to_well', 'cartridge_1', 1, 20, 100,2, 100,10, 20 , 100, 2, 100, 10, 2);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('4fa5c4e3-699c-47bb-ac7a-b26d04efaeb5','4fa5c4e3-699c-47bb-ac7a-b26d04efaeb5', 'well_to_shaker', 'cartridge_1', 1, 20, 100,2, 100,10, 40 , 100, 2, 100, 10, 0);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('fb88bada-ced7-4fa2-b845-4bb91e74341e','fb88bada-ced7-4fa2-b845-4bb91e74341e', 'shaker_to_well', 'cartridge_1', 0, 40, 100,2, 100,10, 20 , 100, 2, 100, 10, 2);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('38a20548-85bf-4c3b-82bc-4d87b87f2dbe','38a20548-85bf-4c3b-82bc-4d87b87f2dbe', 'deck_to_well', 'cartridge_1', 4, 40, 100,2, 100,10, 20 , 100, 2, 100, 10, 2);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('3b47f63e-fb02-460b-864f-edc4a302af5a','3b47f63e-fb02-460b-864f-edc4a302af5a', 'well_to_deck', 'cartridge_1', 1, 20, 100,2, 100,10, 40 , 100, 2, 100, 10, 4);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('8207aa73-74b1-4bca-86ae-88e843ef1eac','8207aa73-74b1-4bca-86ae-88e843ef1eac', 'deck_to_deck', 'cartridge_1', 7, 40, 100,2, 100,10, 40 , 100, 2, 100, 10, 7);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('8207aa73-74b1-4bca-86ae-88e843ef1ea3','a3057838-d5e2-4ac2-9e4b-7d1e4fefd767', 'shaker_to_deck', 'cartridge_1', 7, 40, 100,2, 100,10, 10 , 100, 2, 100, 10, 9);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('8207aa73-74b1-4bca-86ae-88e843ef1ea4','a3057838-d5e2-4ac2-9e4b-7d1e4fefd768', 'deck_to_shaker', 'cartridge_1', 7, 10, 100,2, 100,10, 40 , 100, 2, 100, 10, 7);
--  For cartridge_2
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('bbe172f0-e315-4ebd-83ab-c1c0b531e2f8','bbe172f0-e315-4ebd-83ab-c1c0b531e2f8', 'well_to_well', 'cartridge_2', 1, 20, 100,2, 100,10, 20 , 100, 2, 100, 10, 2);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('cd446adb-f7ed-4368-aa65-b0e04fdd2c81','cd446adb-f7ed-4368-aa65-b0e04fdd2c81', 'well_to_shaker', 'cartridge_2', 1, 20, 100,2, 100,10, 40 , 100, 2, 100, 10, 0);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('e6e7f15e-a06d-4d04-ac60-4de113fd8c42','e6e7f15e-a06d-4d04-ac60-4de113fd8c42', 'shaker_to_well', 'cartridge_2', 0, 40, 100,2, 100,10, 20 , 100, 2, 100, 10, 2);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('62c1afc8-87cb-4bb6-833e-b5716cbf71cd','62c1afc8-87cb-4bb6-833e-b5716cbf71cd', 'deck_to_well', 'cartridge_2', 4, 40, 100,2, 100,10, 20 , 100, 2, 100, 10, 2);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('a3b7da1c-a402-4b81-9e76-eeac4782b779','a3b7da1c-a402-4b81-9e76-eeac4782b779', 'well_to_deck', 'cartridge_2', 1, 20, 100,2, 100,10, 40 , 100, 2, 100, 10, 4);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd766','a3057838-d5e2-4ac2-9e4b-7d1e4fefd766', 'deck_to_deck', 'cartridge_2', 7, 40, 100,2, 100,10, 40 , 100, 2, 100, 10, 7);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('8207aa73-74b1-4bca-86ae-88e843ef1ea1','a3057838-d5e2-4ac2-9e4b-7d1e4fefd769', 'shaker_to_deck', 'cartridge_2', 7, 40, 100,2, 100,10, 10 , 100, 2, 100, 10, 9);
INSERT INTO aspire_dispense (id, process_id,category,cartridge_type,source_position,aspire_height,aspire_mixing_volume,aspire_no_of_cycles,aspire_volume,aspire_air_volume,dispense_height,dispense_mixing_volume,dispense_no_of_cycles,dispense_volume,dispense_blow_volume,destination_position)VALUES ('8207aa73-74b1-4bca-86ae-88e843ef1ea2','a3057838-d5e2-4ac2-9e4b-7d1e4fefd770', 'deck_to_shaker', 'cartridge_2', 7, 10, 100,2, 100,10, 40 , 100, 2, 100, 10, 7);

-- Processes for TipDocking
INSERT INTO tip_dock (id, type,position,height,process_id) VALUES ('bbe172f0-e315-4ebd-83ab-c1c0b531e2f9','deck',8,10,'a3057838-d5e2-4ac2-9e4b-7d1e4fefd769');
INSERT INTO tip_dock (id, type,position,height,process_id) VALUES ('bbe172f0-e315-4ebd-83ab-c1c0b531e2f8','cartridge_1',4,12,'a3057838-d5e2-4ac2-9e4b-7d1e4fefd779');
INSERT INTO tip_dock (id, type,position,height,process_id) VALUES ('bbe172f0-e315-4ebd-83ab-c1c0b531e2f7','cartridge_2',5,11,'a3057838-d5e2-4ac2-9e4b-7d1e4fefd789');






