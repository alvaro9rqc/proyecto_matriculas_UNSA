-- name: CreateStudent :exec
INSERT INTO student (
    code, account_id
) VALUES (
    $1, $2
);

-- name: ListStudents :many
SELECT
    s.id,
    s.code,
    a.email,
    a.name,
    a.surname,
    sg.name AS student_group_name
FROM student s
JOIN account a ON s.account_id = a.id
JOIN student_group sg ON s.student_group_id = sg.id
ORDER BY s.code
LIMIT $1
OFFSET $2;

-- name: FullListStudents :many
SELECT
    s.id,
    s.code,
    a.email,
    a.name,
    a.surname,
    sg.name AS student_group_name
FROM student s
JOIN account a ON s.account_id = a.id
JOIN student_group sg ON s.student_group_id = sg.id
ORDER BY s.code;
