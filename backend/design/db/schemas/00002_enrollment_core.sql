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

CREATE TABLE course (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    credits SMALLINT NOT NULL,
    cicle_number SMALLINT NOT NULL,
    CONSTRAINT credits_value CHECK (credits > 0),
    CONSTRAINT cicle_number_value CHECK (cicle_number > 0)
);

CREATE TABLE major (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) UNIQUE NOT NULL
);

CREATE TABLE modality (
    id SERIAL PRIMARY KEY,
    name VARCHAR(40) UNIQUE NOT NULL
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

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS modality;

DROP TABLE IF EXISTS major;

DROP TABLE IF EXISTS course;

DROP TABLE IF EXISTS speaker;

DROP TABLE IF EXISTS student;

DROP TABLE IF EXISTS installation;

DROP TABLE IF EXISTS student_group;

-- +goose StatementEnd
