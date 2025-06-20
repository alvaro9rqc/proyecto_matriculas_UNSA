-- name: CreateCourse :exec
INSERT INTO course (
    name, credits, cycle_number
) VALUES (
    $1, $2, $3
);

-- name: UpdateCourse :exec
UPDATE course
SET
    name = $1,
    credits = $2,
    cycle_number = $3
WHERE id = $4;

-- name: DeleteCourse :exec
DELETE FROM course
WHERE id = $1;

-- name: ListCourses :many
SELECT
    id,
    name,
    credits,
    cycle_number
FROM course
ORDER BY cycle_number, name;
