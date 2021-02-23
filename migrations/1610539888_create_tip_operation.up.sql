CREATE TYPE tip_operation_type AS ENUM ('pickup', 'discard');

CREATE TABLE  tip_operation(
 id uuid PRIMARY KEY,
 type tip_operation_type DEFAULT 'discard',
 position int CHECK (position<11),
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (id) REFERENCES processes(id) ON UPDATE CASCADE ON DELETE CASCADE);
