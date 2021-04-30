CREATE TABLE IF NOT EXISTS user_auths(
    auth_id uuid default uuid_generate_v4(),
    username text,
    unique(auth_id,username),
    FOREIGN KEY (username) REFERENCES users(username) ON UPDATE CASCADE ON DELETE CASCADE);
);
