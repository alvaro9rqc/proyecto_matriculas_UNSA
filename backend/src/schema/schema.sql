
-- TODO: add role attribute
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(40) UNIQUE NOT NULL, 
  first_name VARCHAR(30) NOT NULL,
  remaining_names VARCHAR(128),
  last_names VARCHAR(30) NOT NULL,
  CONSTRAINT email_format CHECK (email ~* '^[A-Za-z0-9._]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

CREATE TABLE installations (
  id SERIAL PRIMARY KEY,
  name VARCHAR(40) NOT NULL,
  description VARCHAR(255)
);

CREATE TABLE students (
  id SERIAL PRIMARY KEY,
  student_code VARCHAR(30) UNIQUE NOT NULL,
  user_id INTEGER UNIQUE NOT NULL,
  FOREIGN KEY (user_id) 
  REFERENCES users(id)
  ON DELETE CASCADE
);

CREATE TABLE professors (
  id SERIAL PRIMARY KEY,
  user_id INTEGER UNIQUE NOT NULL,
  FOREIGN KEY (user_id) 
  REFERENCES users(id) 
  ON DELETE CASCADE
);

CREATE TABLE semesters (
  year INTEGER,
  number INTEGER,
  PRIMARY KEY(year, number)
);
