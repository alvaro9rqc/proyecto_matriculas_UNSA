-- name: CreateCourse :exec
INSERT INTO course (
    name, credits, cicle_number
) VALUES (
    $1, $2, $3
);

-- name: UpdateCourse :exec
UPDATE course
SET
    name = $1,
    credits = $2,
    cicle_number = $3
WHERE id = $4;

-- name: DeleteCourse :exec
DELETE FROM course
WHERE id = $1;

-- name: ListCourses :many
SELECT
    id,
    name,
    credits,
    cicle_number
FROM course
ORDER BY cicle_number, name;
