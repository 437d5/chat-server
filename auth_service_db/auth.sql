CREATE TABLE IF NOT EXISTS "users" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    password VARCHAR(32) 
);

INSERT INTO users (name, password) VALUES ('hello', '5d41402abc4b2a76b9719d911017c592');
INSERT INTO users (name, password) VALUES ('bye', 'bfa99df33b137bc8fb5f5407d7e58da8');