-- +goose Up
-- +goose StatementBegin

-- Enrollment Major Process Table: (1, 2025, 1, true)
CREATE TABLE enrollment_major_process (
    id SERIAL PRIMARY KEY,
    major_id INTEGER,
    year SMALLINT NOT NULL,
    cicle_type SMALLINT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    FOREIGN KEY (major_id) REFERENCES major (id) ON DELETE CASCADE,
    CONSTRAINT cicle_format CHECK (cicle_type > 0),
    CONSTRAINT year_range CHECK (year > 2024),
    UNIQUE (major_id, year, cicle_type)
);

-- Installation Table: se refiere a los cupos disponibles para un curso
CREATE TABLE tuition (
    id SERIAL PRIMARY KEY,
    total_places INTEGER,
    taken_places INTEGER,
    enrollment_major_process_id INTEGER,
    FOREIGN KEY (enrollment_major_process_id) REFERENCES enrollment_major_process (id)
);

-- Event Table: A los horarios de un curso
CREATE TABLE event (
    id SERIAL PRIMARY KEY,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    tuition_id INTEGER,
    installation_id INTEGER,
    FOREIGN KEY (tuition_id) REFERENCES tuition (id) ON DELETE CASCADE,
    FOREIGN KEY (installation_id) REFERENCES installation (id) ON DELETE CASCADE
);

-- Student Major Table: Se refiere a la carrera profesional del estudiante
CREATE TABLE student_major (
    student_id INTEGER,
    major_id INTEGER,
    PRIMARY KEY (major_id, student_id),
    FOREIGN KEY (major_id) REFERENCES major (id) ON DELETE CASCADE,
    FOREIGN KEY (student_id) REFERENCES student (id) ON DELETE CASCADE
);

-- Student Course Table: Se refiere al curso que un estudiante ha llevado con anterioridad, el numero de intentos y si aprobó o no
CREATE TABLE student_course (
    student_id INTEGER,
    course_id INTEGER,
    attemps SMALLINT NOT NULL,
    passed BOOLEAN NOT NULL,
    PRIMARY KEY (student_id, course_id),
    FOREIGN KEY (student_id) REFERENCES student (id) ON DELETE CASCADE,
    FOREIGN KEY (course_id) REFERENCES course (id)
);

-- Course Major Table: Se refiere a los cursos que son parte de una carrera profesional
CREATE TABLE course_major (
    course_id INTEGER,
    major_id INTEGER,
    PRIMARY KEY (course_id, major_id),
    FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE,
    FOREIGN KEY (major_id) REFERENCES major (id) ON DELETE CASCADE
);

-- Course Prerequisite Table: Se refiere a los cursos que son prerequisitos de otro curso
CREATE TABLE course_prerequisite (
    course_id INTEGER,
    prerequisite_id INTEGER,
    PRIMARY KEY (course_id, prerequisite_id),
    FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE,
    FOREIGN KEY (prerequisite_id) REFERENCES course (id) ON DELETE CASCADE,
    CHECK (course_id <> prerequisite_id)
);

-- Course Modality Table: Se refiere a los cursos que tienen una modalidad especifica y el numero de horas que se dictan
CREATE TABLE course_modality (
    course_id INTEGER,
    modality_id INTEGER,
    hours SMALLINT NOT NULL,
    PRIMARY KEY (course_id, modality_id),
    FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE,
    FOREIGN KEY (modality_id) REFERENCES modality (id) ON DELETE CASCADE,
    CONSTRAINT hours_format CHECK (hours > 0)
);

-- Tuition Speaker Table: Se refiere a los oradores que dictan un curso
CREATE TABLE tuition_speaker (
    tuition_id INTEGER,
    speaker_id INTEGER,
    FOREIGN KEY (tuition_id) REFERENCES tuition (id) ON DELETE CASCADE,
    FOREIGN KEY (speaker_id) REFERENCES speaker (id) ON DELETE CASCADE
);

-- Tuition Modality Course Table: Se refiere a los cursos que se dictan en una modalidad especifica
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

DROP TABLE IF EXISTS enrollment_major_process;

-- +goose StatementEnd
