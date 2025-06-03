-- name: EnrollmentStudentCourse :exec
INSERT INTO
    student_course (student_id, course_id, attemps, passed)
VALUES
    ($1, $2, 1, $3);

-- name: UpdateEnrollmentStudentCourse :exec
UPDATE student_course
SET
    attemps = attemps + 1,
    passed = $3
WHERE
    student_id = $1 AND course_id = $2;

-- name: GetEnrolledUsersByCourse :many
SELECT
    au.first_name,
    au.remaining_names,
    au.last_names,
    au.email
FROM student_course sc
JOIN student s ON sc.student_id = s.id
JOIN account_user au ON s.account_user_id = au.id
WHERE sc.course_id = $1;

-- name: DeleteStudentCourse :exec
DELETE FROM student_course
WHERE student_id = $1 AND course_id = $2;
