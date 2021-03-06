CREATE TYPE magnet_operation_type  AS ENUM ('attach','detach');
-- NOTE: Subtypes are unused after the requirement analysis from Bio team
-- CREATE TYPE magnet_operation_subtype AS ENUM ('lysis', 'wash', 'elusion','full_detach', 'semi_detach');


CREATE TABLE IF NOT EXISTS attach_detach(
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 operation magnet_operation_type,
 height int,
 process_id uuid UNIQUE NOT NULL,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (process_id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
