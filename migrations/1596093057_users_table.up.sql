CREATE TYPE role_type AS  ENUM('admin', 'operator', 'engineer');

CREATE TABLE IF NOT EXISTS users(
    username varchar(50) PRIMARY KEY,
    password varchar(50) NOT NULL,
    role role_type,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP);
