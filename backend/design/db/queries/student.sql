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
    a.surname
FROM student s
JOIN account a ON s.account_id = a.id
ORDER BY a.surname, a.name
LIMIT $1
OFFSET $2;

-- name: FullListStudents :many
SELECT
    s.id,
    s.code,
    a.email,
    a.name,
    a.surname
FROM student s
JOIN account a ON s.account_id = a.id
ORDER BY a.surname, a.name;
