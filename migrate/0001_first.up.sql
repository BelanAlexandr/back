CREATE TABLE roledictionary (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL UNIQUE
);
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) UNIQUE NOT NULL,
    pass VARCHAR(255) NOT NULL,
    role INTEGER NOT NULL REFERENCES roledictionary(id) ON DELETE RESTRICT
);
INSERT INTO roledictionary (id, role_name) VALUES 
(1, 'admin'),
(2, 'director'),
(3, 'employee')
ON CONFLICT (id) DO NOTHING;