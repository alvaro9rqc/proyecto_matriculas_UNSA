-- name: CreateCourse :exec
INSERT INTO course (
    name, credits, cycle_number, process_id
) VALUES (
    $1, $2, $3, $4
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
    *
FROM course
ORDER BY cycle_number, name;
