CREATE TYPE IF NOT EXISTS magnet_operation_type  AS ENUM ('attach','detach');
CREATE TYPE IF NOT EXISTS magnet_operation_subtype AS ENUM ('lysis', 'wash', 'illusion','full_detach', 'semi_detach');


CREATE TABLE IF NOT EXISTS attach_detach(
 id uuid,
 operation magnet_operation_type,
 operation_type magnet_operation_subtype,
 process_id uuid,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (id),
 FOREIGN KEY (process_id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
