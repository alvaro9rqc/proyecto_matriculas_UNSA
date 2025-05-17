-- test for user
INSERT INTO users (email, first_name, remaining_names, last_names)
VALUES ('alvaro9rqc@gmail.com', 'Álvaro',  'Raul', 'Quispe Condori');
--SELECT * FROM users;

-- test for installations
INSERT INTO installations(name, description)
VALUES ('edicio', 'es  de color verde,e stá ubicado en tal tal tal');
--SELECT * FROM installations;


INSERT INTO students (code, user_id)
VALUES ( 
    '20230467',
    (SELECT id FROM users WHERE email = 'alvaro9rqc@gmail.com')
);
-- SELECT * FROM students;

-- test for professors
INSERT INTO users (email, first_name, remaining_names, last_names)
VALUES ('apaz@unsa.edu.pe', 'Alfredo',  '', 'Paz Valderrama');
INSERT INTO professors (user_id)
VALUES ((SELECT id FROM users WHERE email = 'apaz@unsa.edu.pe'));
--SELECT * FROM professors;

INSERT INTO semesters (year, number)
VALUES (2025, 1);
--SELECT * FROM semesters;

INSERT INTO courses (name, credits)
VALUES ('Fundamentos de la programación', 3);
-- SELECT * FROM courses;

INSERT INTO programs (name)
VALUES ('Ingeniería de sistemas');
-- SELECT * FROM programs;
