-- Insert into recipe
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2de4','covid', 'Recipe for covid extraction', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2de5','tip_docking', 'Recipe for tip_docking position on deck', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2de6','tip_docking', 'Recipe for tip_docking position on cartridge 1', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2de7','tip_docking', 'Recipe for tip_docking position on cartridge 2', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2d11','tip 1', 'Recipe for maleria 1 extraction', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2d12','tip 2', 'Recipe for maleria 2 extraction', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2d13','tip 3', 'Recipe for maleria 3 extraction', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2d14','piercing all wells cartridge_1', 'piercing 8 well cartridge all wells', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2d15','piercing only even wells cartridge_1', 'piercing 8 well cartridge only even wells', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2d16','piercing only odd wells cartridge_1', 'piercing 8 well cartridge only odd wells', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2d17','piercing all wells cartridge_2', 'piercing 4 well cartridge all wells', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2d18','piercing only even wells cartridge_2', 'piercing 4 well cartridge only even wells', 1, 2, 3, 4, 5, 1, 6, 2, 7);
INSERT INTO recipes(id, name, description, pos_1, pos_2, pos_3, pos_4, pos_5, pos_cartridge_1, pos_7, pos_cartridge_2, pos_9) VALUES('6b7fcfa2-8337-4d79-829a-e9bd486a2d19','piercing only odd wells cartridge_2', 'piercing 4 well cartridge only odd wells', 1, 2, 3, 4, 5, 1, 6, 2, 7);

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
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('2fa97b44-f3c3-460a-9585-9932b2de1ac2','Heating', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 17, 'HT-FT-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('42479aae-7342-43af-a9d3-520fbffc0f24','Heating', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 18, 'HT-FF-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('fee69b9e-0898-4078-bc28-655ebddbfb5b','Heating', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 19, 'HT-FT-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('1d3334e5-5779-423d-83d5-b8724e7213cb','Heating', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 20, 'HT-FF-5' );

--Delay processes

INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('1d3334e5-5779-423d-83d5-b8724e5213cb','Delay', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 21, 'dl-50' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('1d3334e5-5779-423d-83d5-b8524e5214cb','Delay', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 22, 'dl-60' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('1d3334e5-5779-423d-83d5-b8424e5215cb','Delay', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 23, 'dl-70' );



INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd068','AttachDetach', '6b7fcfa2-8337-4d79-829a-e9bd486a2de8', 20, 'ATDT-AT-1' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd968','AttachDetach', '6b7fcfa2-8337-4d79-829a-e9bd486a2de8', 21, 'ATDT-DT-1' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd769','TipDocking', '6b7fcfa2-8337-4d79-829a-e9bd486a2de5', 9, 'TD-DECK-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd779','TipDocking', '6b7fcfa2-8337-4d79-829a-e9bd486a2de6', 10, 'TD-C1-4' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd789','TipDocking', '6b7fcfa2-8337-4d79-829a-e9bd486a2de7', 11, 'TD-C1-5' );


INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd767','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 7, 'AD-SD-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd768','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 8, 'AD-DS-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('2557f792-e60f-4f91-a79e-575349b5b1e5','Shaking', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 9, 'SH-NH-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('8fc9e765-3120-4617-9085-e9b81d589030','Shaking', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 10, 'SH-WH-NF-6' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('4cc37963-c564-43d4-b277-7c58ef5a0dc7','Shaking', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 11, 'SH-WH-FT-7' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('6543f226-098d-4f35-afd0-ab692382924c','Shaking', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 12, 'SH-NH-8' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('bbd4c820-1eab-4e3f-a508-4b986bc5227b','Shaking', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 13, 'SH-NH-9' );

--  For cartridge_2
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('bbe172f0-e315-4ebd-83ab-c1c0b531e2f8','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 9, 'AD-WW-c2-1-2' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('cd446adb-f7ed-4368-aa65-b0e04fdd2c81','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 10, 'AD-WS-c2-1' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('e6e7f15e-a06d-4d04-ac60-4de113fd8c42','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 11, 'AD-SW-c2-2' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('62c1afc8-87cb-4bb6-833e-b5716cbf71cd','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 12, 'AD-DW-c2-4-2' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3b7da1c-a402-4b81-9e76-eeac4782b779','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 13, 'AD-WD-c2-2-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd766','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 14, 'AD-DD-4-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd769','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 15, 'AD-SD-7' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd770','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 16, 'AD-DS-7' );

-- Insert into processes 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a5d058e3-7ce3-4a42-b2da-690e47139741','TipOperation', '6b7fcfa2-8337-4d79-829a-e9bd486a2d11', 1, 'TO-PK-1' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a5d058e3-7ce3-4a42-b2da-690e47139731','TipOperation', '6b7fcfa2-8337-4d79-829a-e9bd486a2d11', 2, 'TO-DK' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a5d058e3-7ce3-4a42-b2da-690e47139742','TipOperation', '6b7fcfa2-8337-4d79-829a-e9bd486a2d12', 1, 'TO-PK-2' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a5d058e3-7ce3-4a42-b2da-690e47139732','TipOperation', '6b7fcfa2-8337-4d79-829a-e9bd486a2d12', 2, 'TO-DK' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a5d058e3-7ce3-4a42-b2da-690e47139743','TipOperation', '6b7fcfa2-8337-4d79-829a-e9bd486a2d13', 1, 'TO-PK-3' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a5d058e3-7ce3-4a42-b2da-690e47139733','TipOperation', '6b7fcfa2-8337-4d79-829a-e9bd486a2d13', 2, 'TO-DK' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd769','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 15, 'AD-SD-5' );
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a3057838-d5e2-4ac2-9e4b-7d1e4fefd770','AspireDispense', '6b7fcfa2-8337-4d79-829a-e9bd486a2de4', 16, 'AD-DS-5' );

-- For piercing processes
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a6d058e3-7ce3-4a42-b2da-690e37139733','Piercing', '6b7fcfa2-8337-4d79-829a-e9bd486a2d14', 1, 'PI-C1-ALL' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a6d058e3-7ce3-4a42-b2da-690e37139734','Piercing', '6b7fcfa2-8337-4d79-829a-e9bd486a2d15', 1, 'PI-C1-2-4-6-8-0' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a6d058e3-7ce3-4a42-b2da-690e37139735','Piercing', '6b7fcfa2-8337-4d79-829a-e9bd486a2d16', 1, 'PI-C1-1-3-5-7-9' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a6d058e3-7ce3-4a42-b2da-690e37139736','Piercing', '6b7fcfa2-8337-4d79-829a-e9bd486a2d17', 1, 'PI-C2-ALL' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a6d058e3-7ce3-4a42-b2da-690e37139737','Piercing', '6b7fcfa2-8337-4d79-829a-e9bd486a2d18', 1, 'PI-C2-2-4-6-8-0' ); 
INSERT INTO processes(id, type, recipe_id, sequence_num, name ) VALUES ('a6d058e3-7ce3-4a42-b2da-690e37139738','Piercing', '6b7fcfa2-8337-4d79-829a-e9bd486a2d19', 1, 'PI-C2-1-3-5-7-9' ); 

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
INSERT INTO tip_docking (id, type,position,height,process_id) VALUES ('bbe172f0-e315-4ebd-83ab-c1c0b531e2f9','deck',8,10,'a3057838-d5e2-4ac2-9e4b-7d1e4fefd769');
INSERT INTO tip_docking (id, type,position,height,process_id) VALUES ('bbe172f0-e315-4ebd-83ab-c1c0b531e2f8','cartridge_1',4,12,'a3057838-d5e2-4ac2-9e4b-7d1e4fefd779');
INSERT INTO tip_docking (id, type,position,height,process_id) VALUES ('bbe172f0-e315-4ebd-83ab-c1c0b531e2f7','cartridge_2',5,11,'a3057838-d5e2-4ac2-9e4b-7d1e4fefd789');

-- For Heating

INSERT INTO heating(id,process_id,temperature,follow_temp,duration) VALUES('659057e1-8a41-4f2e-a94e-1ffc60aea5a6','2fa97b44-f3c3-460a-9585-9932b2de1ac2',20,true,'20');
INSERT INTO heating(id,process_id,temperature,follow_temp,duration) VALUES('1cdb0a53-f151-4b37-8d8d-8741a044150c','42479aae-7342-43af-a9d3-520fbffc0f24',20,false,'20');
INSERT INTO heating(id,process_id,temperature,follow_temp,duration) VALUES('bee80d98-098a-4677-b3fb-6932278231b8','fee69b9e-0898-4078-bc28-655ebddbfb5b',20,true,'20');
INSERT INTO heating(id,process_id,temperature,follow_temp,duration) VALUES('e0e974f4-2269-4b79-a5be-842be95f02bb','1d3334e5-5779-423d-83d5-b8724e7213cb',20,false,'20');

-- For Tip Operation
INSERT INTO tip_operation (id, process_id, type, position) VALUES ('9207aa73-74b1-4bca-86ae-88e843ef1eaa','a5d058e3-7ce3-4a42-b2da-690e47139741', 'pickup', 1);
INSERT INTO tip_operation (id, process_id, type, position) VALUES ('9207aa73-74b1-4bca-86ae-88e843ef1e1a','a5d058e3-7ce3-4a42-b2da-690e47139731', 'discard', 0);
INSERT INTO tip_operation (id, process_id, type, position) VALUES ('a207aa73-74b1-4bca-86ae-88e843ef1eab','a5d058e3-7ce3-4a42-b2da-690e47139742', 'pickup', 2);
INSERT INTO tip_operation (id, process_id, type, position) VALUES ('a207aa73-74b1-4bca-86ae-88e843ef1e1b','a5d058e3-7ce3-4a42-b2da-690e47139732', 'discard', 0);
INSERT INTO tip_operation (id, process_id, type, position) VALUES ('b207aa73-74b1-4bca-86ae-88e843ef1eac','a5d058e3-7ce3-4a42-b2da-690e47139743', 'pickup', 3);
INSERT INTO tip_operation (id, process_id, type, position) VALUES ('b207aa73-74b1-4bca-86ae-88e843ef1e1c','a5d058e3-7ce3-4a42-b2da-690e47139733', 'discard', 0);

-- attach Detach
INSERT INTO attach_detach (id,operation,operation_type,process_id) VALUES('16587ff5-5d49-4ab7-b155-59069a380ff7','attach','wash','a3057838-d5e2-4ac2-9e4b-7d1e4fefd068');
INSERT INTO attach_detach (id,operation,operation_type,process_id) VALUES('16587ff5-5d49-4ab7-b155-59069a380ff8','detach','full_detach','a3057838-d5e2-4ac2-9e4b-7d1e4fefd968');
-- For Delay
INSERT INTO delay (id, delay_time, process_id) VALUES ('9207aa73-74b1-4bca-85ae-88e843ef1eaa',50,'1d3334e5-5779-423d-83d5-b8724e5213cb');
INSERT INTO delay (id, delay_time, process_id) VALUES ('9207aa73-74b1-4bca-85ae-88e843ef1eaa',60,'1d3334e5-5779-423d-83d5-b8524e5214cb');
INSERT INTO delay (id, delay_time, process_id) VALUES ('9207aa73-74b1-4bca-85ae-88e843ef1eaa',70,'1d3334e5-5779-423d-83d5-b8424e5215cb');


-- For Piercing
INSERT INTO piercing (id,type,cartridge_wells,discard,process_id) VALUES ('7a7e3565-bdfd-4a2d-9b45-f2147d33c082','cartridge_1','{1,2,3,4,5,6,7,8}', 'at_discard_box','a6d058e3-7ce3-4a42-b2da-690e37139733');
INSERT INTO piercing (id,type,cartridge_wells,discard,process_id) VALUES ('7a7e3565-bdfd-4a2d-9b45-f2147d33c083','cartridge_1','{2,4,6,8}', 'at_discard_box','a6d058e3-7ce3-4a42-b2da-690e37139734');
INSERT INTO piercing (id,type,cartridge_wells,discard,process_id) VALUES ('7a7e3565-bdfd-4a2d-9b45-f2147d33c084','cartridge_1','{1,3,5,7}', 'at_discard_box','a6d058e3-7ce3-4a42-b2da-690e37139735');
INSERT INTO piercing (id,type,cartridge_wells,discard,process_id) VALUES ('7a7e3565-bdfd-4a2d-9b45-f2147d33c085','cartridge_2','{1,2,3,4}', 'at_discard_box','a6d058e3-7ce3-4a42-b2da-690e37139736');
INSERT INTO piercing (id,type,cartridge_wells,discard,process_id) VALUES ('7a7e3565-bdfd-4a2d-9b45-f2147d33c086','cartridge_2','{2,4}', 'at_discard_box','a6d058e3-7ce3-4a42-b2da-690e37139737');
INSERT INTO piercing (id,type,cartridge_wells,discard,process_id) VALUES ('7a7e3565-bdfd-4a2d-9b45-f2147d33c087','cartridge_2','{1,3}', 'at_discard_box','a6d058e3-7ce3-4a42-b2da-690e37139738');
--Insert into shaking
INSERT INTO shaking (id,process_id,with_temp,follow_temp,temperature,rpm_1,rpm_2,time_1,time_2) VALUES('bbe172f0-e315-4ebd-83ab-c1c0b531e2f6','2557f792-e60f-4f91-a79e-575349b5b1e5',false,false,0,500,800,60,60);
INSERT INTO shaking (id,process_id,with_temp,follow_temp,temperature,rpm_1,rpm_2,time_1,time_2) VALUES('a03fc110-5b9e-4461-b114-dfd21bae77d8','8fc9e765-3120-4617-9085-e9b81d589030',true,false,0,500,800,120,120);
INSERT INTO shaking (id,process_id,with_temp,follow_temp,temperature,rpm_1,rpm_2,time_1,time_2) VALUES('2066b0aa-7688-4fa7-bba1-c5836a3ab01b','4cc37963-c564-43d4-b277-7c58ef5a0dc7',true,true,600,500,800,180,180);
INSERT INTO shaking (id,process_id,with_temp,follow_temp,temperature,rpm_1,rpm_2,time_1,time_2) VALUES('c6d62e6c-1c19-4d3d-862e-a8dc5f8c2629','6543f226-098d-4f35-afd0-ab692382924c',false,false,0,6500,14500,240,240);
INSERT INTO shaking (id,process_id,with_temp,follow_temp,temperature,rpm_1,rpm_2,time_1,time_2) VALUES('0c911dfc-2993-4992-8fa6-e6d779e8814b','bbd4c820-1eab-4e3f-a508-4b986bc5227b',false,false,0,500,800,300,300);
