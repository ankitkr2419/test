CREATE TYPE tip_operation_type AS ENUM ('pickup', 'discard');

CREATE TABLE  tip_operation(
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 type tip_operation_type DEFAULT 'discard',
 position int CHECK (position<11),
 process_id uuid,
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (process_id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
