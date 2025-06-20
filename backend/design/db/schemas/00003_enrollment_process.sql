-- +goose Up
-- +goose StatementBegin

CREATE TABLE section (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    course_id INTEGER NOT NULL,
    FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE
);

CREATE TABLE slots (
    id SERIAL PRIMARY KEY,
    total_places INTEGER,
    taken_places INTEGER,
    section_id INTEGER NOT NULL,
    FOREIGN KEY (section_id) REFERENCES section (id) ON DELETE CASCADE
);

-- Event Table: A los horarios de un curso
CREATE TABLE event (
    id SERIAL PRIMARY KEY,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    section_id INTEGER NOT NULL,
    installation_id INTEGER,
    modality_id INTEGER NOT NULL,

    FOREIGN KEY (modality_id) REFERENCES modality (id) ON DELETE RESTRICT,
    FOREIGN KEY (section_id) REFERENCES section (id) ON DELETE CASCADE,
    FOREIGN KEY (installation_id) REFERENCES installation (id) ON DELETE CASCADE
);

-- Student Major Table: Se refiere a la carrera profesional del estudiante
CREATE TABLE student_major (
    student_id INTEGER,
    major_id INTEGER,
    PRIMARY KEY (student_id, major_id),
    FOREIGN KEY (major_id) REFERENCES major (id) ON DELETE CASCADE,
    FOREIGN KEY (student_id) REFERENCES student (id) ON DELETE CASCADE
);

-- Student Course Table: Se refiere al curso que un estudiante ha llevado con anterioridad, el numero de intentos y si aprobó o no
CREATE TABLE student_course (
    student_id INTEGER,
    course_id INTEGER,
    attempts SMALLINT NOT NULL,
    passed BOOLEAN NOT NULL,
    PRIMARY KEY (student_id, course_id),
    FOREIGN KEY (student_id) REFERENCES student (id) ON DELETE CASCADE,
    FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE
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

-- Tuition Speaker Table: Se refiere a los oradores que dictan un curso
CREATE TABLE section_speaker (
    section_id INTEGER,
    speaker_id INTEGER,
    FOREIGN KEY (section_id) REFERENCES section (id) ON DELETE CASCADE,
    FOREIGN KEY (speaker_id) REFERENCES speaker (id) ON DELETE CASCADE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS section_speaker;

DROP TABLE IF EXISTS course_prerequisite;

DROP TABLE IF EXISTS student_course;

DROP TABLE IF EXISTS student_major;

DROP TABLE IF EXISTS event;

DROP TABLE IF EXISTS slots;

DROP TABLE IF EXISTS section;

-- +goose StatementEnd
