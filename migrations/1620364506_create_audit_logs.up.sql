
CREATE TYPE audit_activity_type  AS ENUM ('api','db','machine');
CREATE TYPE audit_state_type  AS ENUM ('initialised','completed','aborted','error','paused','resumed');



CREATE TABLE IF NOT EXISTS audit_logs(
    id uuid default uuid_generate_v4(),
    username text NOT NULL,
    activity_type audit_activity_type,
    state_type audit_state_type,
    deck text,
    description text,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (username) REFERENCES users(username)
);

