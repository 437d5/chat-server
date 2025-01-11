CREATE DATABASE auth_db;

\c auth_db;

CREATE TABLE IF NOT EXISTS "users" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    password VARCHAR(32) 
);

INSERT INTO users (name, password) VALUES ('hello', '5d41402abc4b2a76b9719d911017c592');
INSERT INTO users (name, password) VALUES ('bye', 'bfa99df33b137bc8fb5f5407d7e58da8');

CREATE USER auth_service WITH PASSWORD 'auth_service_pass';
GRANT CONNECT ON DATABASE auth_db TO auth_service;
GRANT USAGE ON SCHEMA public TO auth_service;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO auth_service;
GRANT INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO auth_service;
