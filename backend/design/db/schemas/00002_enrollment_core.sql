-- +goose Up
-- +goose StatementBegin
CREATE TABLE student_group (
    id SMALLSERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    priority SMALLINT NOT NULL,
    start_day DATE NOT NULL,
    end_day DATE NOT NULL
);

CREATE TABLE installation (
    id SERIAL PRIMARY KEY,
    name VARCHAR(40) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE student (
    id SERIAL PRIMARY KEY,
    code VARCHAR(30) UNIQUE NOT NULL,
    account_id INTEGER UNIQUE NOT NULL,
    student_group_id SMALLINT,
    FOREIGN KEY (account_id) REFERENCES account (id) ON DELETE CASCADE,
    FOREIGN KEY (student_group_id) REFERENCES student_group (id) ON DELETE RESTRICT
);

CREATE TABLE speaker (
    id SERIAL PRIMARY KEY,
    account_id INTEGER UNIQUE NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account (id) ON DELETE CASCADE
);

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
    FOREIGN KEY (program_id) REFERENCES program (id) ON DELETE CASCADE,
    CONSTRAINT cicle_format CHECK (cicle_type > 0),
    CONSTRAINT year_range CHECK (year > 2024),
    UNIQUE (program_id, year, cicle_type)
);

CREATE TABLE tuition (
    id SERIAL PRIMARY KEY,
    total_places INTEGER,
    taken_places INTEGER,
    enrollment_process_id INTEGER,
    FOREIGN KEY (enrollment_process_id) REFERENCES enrollment_process (id)
);

CREATE TABLE event (
    id SERIAL PRIMARY KEY,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    tuition_id INTEGER,
    installation_id INTEGER,
    FOREIGN KEY (tuition_id) REFERENCES tuition (id) ON DELETE CASCADE,
    FOREIGN KEY (installation_id) REFERENCES installation (id) ON DELETE CASCADE
);

CREATE TABLE student_program (
    student_id INTEGER,
    program_id INTEGER,
    PRIMARY KEY (program_id, student_id),
    FOREIGN KEY (program_id) REFERENCES program (id) ON DELETE CASCADE,
    FOREIGN KEY (student_id) REFERENCES student (id) ON DELETE CASCADE
);

CREATE TABLE student_course (
    student_id INTEGER,
    course_id INTEGER,
    attemps SMALLINT NOT NULL,
    passed BOOLEAN NOT NULL,
    PRIMARY KEY (student_id, course_id),
    FOREIGN KEY (student_id) REFERENCES student (id) ON DELETE CASCADE,
    FOREIGN KEY (course_id) REFERENCES course (id)
);

CREATE TABLE course_program (
    course_id INTEGER,
    program_id INTEGER,
    PRIMARY KEY (course_id, program_id),
    FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE,
    FOREIGN KEY (program_id) REFERENCES program (id) ON DELETE CASCADE
);

CREATE TABLE course_prerequisite (
    course_id INTEGER,
    prerequisite_id INTEGER,
    PRIMARY KEY (course_id, prerequisite_id),
    FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE,
    FOREIGN KEY (prerequisite_id) REFERENCES course (id) ON DELETE CASCADE,
    CHECK (course_id <> prerequisite_id)
);

CREATE TABLE modality (
    id SERIAL PRIMARY KEY,
    name VARCHAR(40) UNIQUE NOT NULL
);

CREATE TABLE course_modality (
    course_id INTEGER,
    modality_id INTEGER,
    hours SMALLINT NOT NULL,
    PRIMARY KEY (course_id, modality_id),
    FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE,
    FOREIGN KEY (modality_id) REFERENCES modality (id) ON DELETE CASCADE,
    CONSTRAINT hours_format CHECK (hours > 0)
);

CREATE TABLE tuition_speaker (
    tuition_id INTEGER,
    speaker_id INTEGER,
    FOREIGN KEY (tuition_id) REFERENCES tuition (id) ON DELETE CASCADE,
    FOREIGN KEY (speaker_id) REFERENCES speaker (id) ON DELETE CASCADE
);

CREATE TABLE tuition_modality_course (
    tuition_id INTEGER,
    modality_id INTEGER,
    course_id INTEGER,
    PRIMARY KEY (tuition_id, modality_id, course_id),
    FOREIGN KEY (tuition_id) REFERENCES tuition (id) ON DELETE CASCADE,
    FOREIGN KEY (modality_id) REFERENCES modality (id) ON DELETE CASCADE,
    FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tuition_modality_course;

DROP TABLE IF EXISTS tuition_speaker;

DROP TABLE IF EXISTS course_modality;

DROP TABLE IF EXISTS course_prerequisite;

DROP TABLE IF EXISTS course_program;

DROP TABLE IF EXISTS student_course;

DROP TABLE IF EXISTS student_program;

DROP TABLE IF EXISTS event;

DROP TABLE IF EXISTS tuition;

DROP TABLE IF EXISTS enrollment_process;

DROP TABLE IF EXISTS modality;

DROP TABLE IF EXISTS program;

DROP TABLE IF EXISTS course;

DROP TABLE IF EXISTS speaker;

DROP TABLE IF EXISTS student;

DROP TABLE IF EXISTS installation;

DROP TABLE IF EXISTS student_group;

-- +goose StatementEnd
