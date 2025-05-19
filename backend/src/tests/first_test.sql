INSERT INTO student_group (name, priority)
VALUES( 'Primer tercio', 1 ),
      ( 'Segundo tercio', 2 ),
      ( 'Tercer tercio', 3 );
-- SELECT * FROM student_groups;

-- test for user
INSERT INTO account_user (email, first_name, remaining_names, last_names, student_group_id)
VALUES 
      ('alvaro9rqc@gmail.com', 'Álvaro',  'Raul', 'Quispe Condori', 1),
      ('romina@gmail.com', 'Romina',  'Marlene', 'Davila Quispe', 2);
-- SELECT * FROM users;

-- test for installations
INSERT INTO installation(name, description)
VALUES ('edicio', 'es  de color verde,e stá ubicado en tal tal tal');
--SELECT * FROM installations;


INSERT INTO student (code, account_user_id)
VALUES ( 
    '20230467',
    (SELECT id FROM account_user WHERE email = 'alvaro9rqc@gmail.com')
);
-- SELECT * FROM student;

-- test for professors
INSERT INTO account_user (email, first_name, remaining_names, last_names)
VALUES ('apaz@unsa.edu.pe', 'Alfredo',  '', 'Paz Valderrama');
INSERT INTO professor (account_user_id)
VALUES ((SELECT id FROM account_user WHERE email = 'apaz@unsa.edu.pe'));
--SELECT * FROM professors;

INSERT INTO semester (year, number)
VALUES (2025, 1);
--SELECT * FROM semesters;

INSERT INTO course (name, credits)
VALUES ('Fundamentos de la programación', 3);
-- SELECT * FROM courses;

INSERT INTO program (name)
VALUES ('Ingeniería de sistemas');
-- SELECT * FROM program;

INSERT INTO student_program(student_id, program_id)
VALUES (
  -- student_id
  (SELECT id FROM student WHERE account_user_id = (SELECT id FROM account_user WHERE email = 'alvaro9rqc@gmail.com')),
  -- program_id
  (SELECT id FROM program WHERE name = 'Ingeniería de sistemas')
);
SELECT * FROM student_program;
