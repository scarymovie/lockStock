CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       uid VARCHAR(36) NOT NULL UNIQUE,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
