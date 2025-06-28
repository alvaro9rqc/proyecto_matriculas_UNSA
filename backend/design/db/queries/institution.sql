CREATE TABLE institution (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) UNIQUE NOT NULL,
    logo_url VARCHAR(255)  -- Optional logo/avatar URL
);
