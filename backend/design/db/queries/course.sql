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

-- name: ListAllCoursesAvailableByStudentInProcess :many
SELECT
    c.id AS course_id,
    c.name AS course_name,
    c.credits,
    c.cycle_number
FROM course c
JOIN student_avalible_courses sac ON sac.course_id = c.id
WHERE sac.student_id = $1 AND c.process_id = $2;
