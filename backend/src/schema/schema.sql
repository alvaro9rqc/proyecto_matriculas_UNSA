/*
This schema is divided in entities and cross reference tables

- They are some rules:
  - A program from course should be inside a program from the attendee
    - Maybe a trigger can solve this during enrollment process
  - There shouldn't be a cicly in course_prerequisite

*/ 



---------------
-- DOMAIN types
---------------

CREATE DOMAIN email_type AS VARCHAR(40)
CHECK (VALUE ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$');

/*
NOTE: attendee table relies on this, there isn't a attendee without a attendee_group
*/

CREATE TABLE attendee_group (
  id SMALLSERIAL PRIMARY KEY,
  name VARCHAR(30) NOT NULL,
  priority SMALLINT NOT NULL
);

CREATE TABLE account_user (
  id SERIAL PRIMARY KEY,
  email email_type NOT NULL,
  first_name VARCHAR(30) NOT NULL,
  remaining_names VARCHAR(128),
  last_names VARCHAR(30) NOT NULL,
  attendee_group_id SMALLINT,
  FOREIGN KEY (attendee_group_id)
  REFERENCES attendee_group(id) ON DELETE RESTRICT
);


CREATE TABLE installation (
  id SERIAL PRIMARY KEY,
  name VARCHAR(40) NOT NULL,
  description VARCHAR(255)
);

CREATE TABLE attendee (
  id SERIAL PRIMARY KEY,
  code VARCHAR(30) UNIQUE NOT NULL,
  account_user_id INTEGER UNIQUE NOT NULL,
  FOREIGN KEY (account_user_id) 
  REFERENCES account_user(id)
  ON DELETE CASCADE
);

CREATE TABLE speaker (
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

CREATE TABLE enrollment_process (
  id SERIAL PRIMARY KEY,
  program_id INTEGER,
  year SMALLINT NOT NULL,
  cicle_type SMALLINT NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT true,
  FOREIGN KEY (program_id) REFERENCES program(id) ON DELETE CASCADE,
  CONSTRAINT cicle_format CHECK (cicle_type > 0),
  CONSTRAINT year_range CHECK (year > 2024),
  UNIQUE (program_id, year, cicle_type)
);

CREATE TABLE modality(
  id SERIAL PRIMARY KEY,
  name VARCHAR(40) UNIQUE NOT NULL
);

CREATE TABLE tuition (
  id SERIAL PRIMARY KEY,
  total_places INTEGER,
  taken_places INTEGER,
  enrollment_process_id INTEGER,
  FOREIGN KEY (enrollment_process_id) REFERENCES enrollment_process(id)
);

CREATE TABLE event (
  id SERIAL PRIMARY KEY,
  start_day DATE NOT NULL,
  end_day DATE NOT NULL,
  start_time TIME NOT NULL,
  end_time TIME NOT NULL,
  weekday SMALLINT NOT NULL,
  tuition_id INTEGER,
  installation_id INTEGER,
  FOREIGN KEY(tuition_id) REFERENCES tuition(id) ON DELETE CASCADE,
  FOREIGN KEY(installation_id) REFERENCES installation(id) ON DELETE CASCADE
);







-------------------------
-------------------------
-------------------------
-------------------------
-- Cross reference tables
-------------------------
-------------------------
-------------------------
-------------------------

-- TODO: What I should do if a program is deleted

CREATE TABLE attendee_program (
  attendee_id INTEGER,
  program_id INTEGER,
  PRIMARY KEY (program_id, attendee_id),
  FOREIGN KEY (program_id) REFERENCES program(id) ON DELETE CASCADE,
  FOREIGN KEY (attendee_id) REFERENCES attendee(id) ON DELETE CASCADE
);

-- Poblem: un estudiante solo se puede matricular a los cursos que su carrera permita. 
-- debería ser course una entidad debil porque depende también de su carrera?
    -- no, porque el curso es independiente de la carrera.

-- I need to CHECK that a course and attendee have at least one course in common

CREATE TABLE attendee_course(
  attendee_id INTEGER,
  course_id INTEGER,
  attemps SMALLINT NOT NULL,
  passed BOOLEAN NOT NULL,
  PRIMARY KEY (attendee_id, course_id), 
  FOREIGN KEY (attendee_id) REFERENCES attendee(id) ON DELETE CASCADE,
  FOREIGN KEY (course_id) REFERENCES course(id)
);

CREATE TABLE course_program (
  course_id INTEGER,
  program_id INTEGER,
  PRIMARY KEY (course_id, program_id),
  FOREIGN KEY (course_id) REFERENCES course(id) ON DELETE CASCADE,
  FOREIGN KEY (program_id) REFERENCES program(id) ON DELETE CASCADE
);

CREATE TABLE course_prerequisite (
  course_id INTEGER, 
  prerequisite_id INTEGER,
  PRIMARY KEY (course_id, prerequisite_id),
  FOREIGN KEY (course_id) REFERENCES course(id) ON DELETE CASCADE,
  FOREIGN KEY (prerequisite_id) REFERENCES course(id) ON DELETE CASCADE,
  CHECK (course_id <> prerequisite_id)
);


CREATE TABLE course_modality(
  course_id INTEGER,
  modality_id INTEGER,
  hours SMALLINT NOT NULL,
  PRIMARY KEY (course_id, modality_id),
  FOREIGN KEY (course_id) REFERENCES course(id) ON DELETE CASCADE,
  FOREIGN KEY (modality_id) REFERENCES modality(id) ON DELETE CASCADE,
  CONSTRAINT hours_format CHECK(hours > 0)
);

CREATE TABLE tuition_speaker(
  tuition_id INTEGER,
  speaker_id INTEGER,
  FOREIGN KEY (tuition_id) REFERENCES tuition(id) ON DELETE CASCADE,
  FOREIGN KEY (speaker_id) REFERENCES speaker(id) ON DELETE CASCADE
);



/* 
Como identiticar a un turno?
  - Está conformado por varias modalidades de un mismo programa
  - Puede tener misma modalidad
  - Pero tiene diferentes horas

  - La suma de las horas del turno. debe ser igual a la suma de horas de las modalidades que incluye
   xd
*/
