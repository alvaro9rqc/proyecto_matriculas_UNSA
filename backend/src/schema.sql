
-- TODO: add role attribute
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(40) UNIQUE NOT NULL, 
    first_name VARCHAR(30) NOT NULL,
    remaining_names VARCHAR(128),
    last_names VARCHAR(30) NOT NULL,
    CONSTRAINT email_format CHECK (email ~* '^[A-Za-z0-9._]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

-- INSERT INTO users (email, first_name, remaining_names, last_names)
-- VALUES ('alvaro9@gmail.com', 'Álvaro',  'Raul', 'Quispe Condori');
-- SELECT * FROM users;

CREATE TABLE installations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(40) NOT NULL,
    description VARCHAR(255)
);

-- INSERT INTO installations(name, description)
-- VALUES ('edicio', 'es  de color verde,e stá ubicado en tal tal tal');
-- SELECT * FROM installations;

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    student_code VARCHAR(30) UNIQUE NOT NULL,
    user_id INTEGER UNIQUE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- INSERT INTO users (email, first_name, remaining_names, last_names)
-- VALUES ('alvaro9rqc@gmail.com', 'Álvaro',  'Raul', 'Quispe Condori');
-- INSERT INTO students (student_code, user_id)
-- VALUES ( 
--     '20230467',
--     (SELECT id FROM users WHERE email = 'alvaro9rqc@gmail.com')
-- );
-- SELECT * FROM students;
