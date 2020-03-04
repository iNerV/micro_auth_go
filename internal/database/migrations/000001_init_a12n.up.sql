CREATE TABLE IF NOT EXISTS users
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    is_active  bool NOT NULL DEFAULT true,
    is_admin   bool NOT NULL DEFAULT false,
    is_staff   bool NOT NULL DEFAULT false,
    username   varchar(250) NOT NULL UNIQUE,
    email      varchar(1000) NOT NULL UNIQUE,
    password   text
);
