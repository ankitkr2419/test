CREATE TYPE IF NOT EXISTS operation  AS ENUM ('attach','detach');
CREATE TYPE IF NOT EXISTS operation_type AS ENUM ('lysis', 'wash', 'illusion','full_detach', 'semi_detach');


CREATE TABLE IF NOT EXISTS attach_detach(
 id uuid,
 operation operation,
 operation_type operation_type,
 process_id uuid,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (id),
 FOREIGN KEY (process_id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
