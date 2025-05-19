/*
NOTE: student table relies on this, there isn't a student without a student_group
*/
CREATE TABLE student_group (
  id SMALLSERIAL PRIMARY KEY,
  name VARCHAR(30) NOT NULL,
  priority SMALLINT NOT NULL
);

CREATE TABLE account_user (
  id SERIAL PRIMARY KEY,
  email VARCHAR(40) UNIQUE NOT NULL, 
  first_name VARCHAR(30) NOT NULL,
  remaining_names VARCHAR(128),
  last_names VARCHAR(30) NOT NULL,
  student_group_id SMALLINT,
  CONSTRAINT email_format CHECK (email ~* '^[A-Za-z0-9._]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
  FOREIGN KEY (student_group_id)
  REFERENCES student_group(id) ON DELETE RESTRICT
);


CREATE TABLE installation (
  id SERIAL PRIMARY KEY,
  name VARCHAR(40) NOT NULL,
  description VARCHAR(255)
);

CREATE TABLE student (
  id SERIAL PRIMARY KEY,
  code VARCHAR(30) UNIQUE NOT NULL,
  account_user_id INTEGER UNIQUE NOT NULL,
  FOREIGN KEY (account_user_id) 
  REFERENCES account_user(id)
  ON DELETE CASCADE
);

CREATE TABLE professor (
  id SERIAL PRIMARY KEY,
  account_user_id INTEGER UNIQUE NOT NULL,
  FOREIGN KEY (account_user_id) 
  REFERENCES account_user(id) 
  ON DELETE CASCADE
);

-- cicle_number is not requited, just 
CREATE TABLE course (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  credits SMALLINT NOT NULL,
  cicle_number SMALLINT NOT NULL,
  CONSTRAINT credits_value CHECK (credits > 0),
  CONSTRAINT cicle_number_value CHECK (cicle_number > 0)
);

CREATE TABLE program (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) UNIQUE NOT NULL
);

CREATE TABLE enrolment_process (
  program_id INTEGER,
  year SMALLINT NOT NULL,
  cicle_type SMALLINT NOT NULL,
  PRIMARY KEY ( program_id, year, cicle_type ),
  FOREIGN KEY (program_id) REFERENCES program(id) ON DELETE CASCADE,
  CONSTRAINT cicle_format CHECK (cicle_type > 0)
);

--CREATE TABLE depa

-------------------------
-- Cross reference tables
-------------------------

-- TODO: What I should do if a cicle is deleted
CREATE TABLE student_program (
  student_id INTEGER,
  program_id INTEGER,
  PRIMARY KEY (program_id, student_id),
  FOREIGN KEY (program_id) REFERENCES program(id),
  FOREIGN KEY (student_id) REFERENCES student(id)
);

-- Poblem: un estudiante solo se puede matricular a los cursos que su carrera permita. 
-- debería ser course una entidad debil porque depende también de su carrera?
    -- no, porque el curso es independiente de la carrera.

CREATE TABLE student_course(
  student_id INTEGER,
  course_id INTEGER,
  attemps SMALLINT NOT NULL,
  passed BOOLEAN NOT NULL,
  PRIMARY KEY (student_id, course_id), 
  FOREIGN KEY (student_id) REFERENCES student(id) ON DELETE CASCADE,
  FOREIGN KEY (course_id) REFERENCES course(id)
);
