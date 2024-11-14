CREATE TABLE rooms
(
    id         SERIAL PRIMARY KEY,
    uid        VARCHAR(36) NOT NULL UNIQUE,
    name       varchar(36) NOT NULL,
    code       varchar(36) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
