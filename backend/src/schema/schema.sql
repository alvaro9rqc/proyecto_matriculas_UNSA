/*
NOTE: students table relies on this, there isn't a students without a group
*/
CREATE TABLE groups (
  id SMALLSERIAL PRIMARY KEY,
  name VARCHAR(30) NOT NULL,
  priority SMALLINT NOT NULL
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(40) UNIQUE NOT NULL, 
  first_name VARCHAR(30) NOT NULL,
  remaining_names VARCHAR(128),
  last_names VARCHAR(30) NOT NULL,
  group_id SMALLINT,
  CONSTRAINT email_format CHECK (email ~* '^[A-Za-z0-9._]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
  FOREIGN KEY (group_id)
  REFERENCES groups(id) ON DELETE RESTRICT
);


CREATE TABLE installations (
  id SERIAL PRIMARY KEY,
  name VARCHAR(40) NOT NULL,
  description VARCHAR(255)
);

CREATE TABLE students (
  id SERIAL PRIMARY KEY,
  code VARCHAR(30) UNIQUE NOT NULL,
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
  year SMALLINT,
  number SMALLINT,
  PRIMARY KEY(year, number)
);

CREATE TABLE courses (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  credits SMALLINT NOT NULL
);

CREATE TABLE programs (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL
);

-- Cross reference tables
