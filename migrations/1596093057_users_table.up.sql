
CREATE TABLE users
(
    username varchar(250),
    password varchar(250),
    role varchar(250),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
