INSERT INTO attendee_group (name, priority)
VALUES( 'Primer tercio', 1 ),
      ( 'Segundo tercio', 2 ),
      ( 'Tercer tercio', 3 );
-- SELECT * FROM attendee_groups;

-- test for user
INSERT INTO account_user (email, first_name, remaining_names, last_names, attendee_group_id)
VALUES 
      ('alvaro9rqc@gmail.com', 'Álvaro',  'Raul', 'Quispe Condori', 1),
      ('romina@gmail.com', 'Romina',  'Marlene', 'Davila Quispe', 2);
-- SELECT * FROM users;

-- test for installations
INSERT INTO installation(name, description)
VALUES ('edicio', 'es  de color verde,e stá ubicado en tal tal tal');
--SELECT * FROM installations;


INSERT INTO attendee (code, account_user_id)
VALUES ( 
    '20230467',
    (SELECT id FROM account_user WHERE email = 'alvaro9rqc@gmail.com')
);
-- SELECT * FROM attendee;




-- test for speakers
INSERT INTO account_user (email, first_name, remaining_names, last_names)
VALUES ('apaz@unsa.edu.pe', 'Alfredo',  '', 'Paz Valderrama');
INSERT INTO speaker (account_user_id)
VALUES ((SELECT id FROM account_user WHERE email = 'apaz@unsa.edu.pe'));
--SELECT * FROM speakers;

-- INSERT INTO cicle (year, number)
-- VALUES (2025, 1);
-- --SELECT * FROM cicles;




INSERT INTO course (name, credits, cicle_number)
VALUES ('Fundamentos de la programación', 3, 1),
       ('Programación web', 4, 2);
-- SELECT * FROM course;




INSERT INTO program (name)
VALUES ('Ingeniería de sistemas');
-- SELECT * FROM program;




INSERT INTO attendee_program(attendee_id, program_id)
VALUES (
  -- attendee_id
  (SELECT id FROM attendee WHERE account_user_id = (SELECT id FROM account_user WHERE email = 'alvaro9rqc@gmail.com')),
  -- program_id
  (SELECT id FROM program WHERE name = 'Ingeniería de sistemas')
);
--SELECT * FROM attendee_program;



-------------------
--test course_program
-------------------
INSERT INTO course_program(course_id, program_id)
VALUES ( 1, 1 );
--SELECT * FROM course_program;





-------------------
--test course_prerequisite
-------------------
INSERT INTO course_prerequisite(course_id, prerequisite_id)
VALUES (2,1);






-------------------
--test modality
-------------------
INSERT INTO modality(name)
VALUES ('Laboratorio'), ('Teoría'), ('Práctica');






-------------------
--test modality
-------------------
INSERT INTO course_modality(course_id, modality_id, hours)
VALUES
  (1, (SELECT id FROM modality WHERE name = 'Laboratorio'), 2),
  (1, (SELECT id FROM modality WHERE name = 'Teoría'), 3),
  (2, (SELECT id FROM modality WHERE name = 'Laboratorio'), 2),
  (2, (SELECT id FROM modality WHERE name = 'Práctica'), 2);
