-- name: CreateStudentMajor :exec
INSERT INTO student_major (
    student_id, major_id
) VALUES (
    $1, $2
);

-- name: DeleteStudentMajor :exec
DELETE FROM student_major
WHERE student_id = $1 AND major_id = $2;

-- name: ListStudentMajors :many
SELECT
    sm.student_id,
    sm.major_id,
    m.name AS major_name
FROM student_major sm
JOIN major m ON sm.major_id = m.id
WHERE sm.student_id = $1
ORDER BY m.name;

-- name: ListStudentByMajor :many
SELECT
    sm.student_id,
    sm.major_id,
    s.code AS student_code
FROM student_major sm
JOIN student s ON sm.student_id = s.id
WHERE sm.major_id = $1
ORDER BY s.code;

-- name: ListMajorsByStudent :many
SELECT
    m.id AS major_id,
    m.name AS major_name
FROM student_major sm
JOIN major m ON m.id = sm.major_id
WHERE sm.student_id = $1
ORDER BY m.name;
