-- name: EnrollmentAttendeeCourse :exec
INSERT INTO
    attendee_course (attendee_id, course_id, attemps, passed)
VALUES
    ($1, $2, 1, $3);

-- name: UpdateEnrollmentAttendeeCourse :exec
UPDATE attendee_course
SET
    attemps = attemps + 1,
    passed = $3
WHERE
    attendee_id = $1 AND course_id = $2;

-- name: GetEnrolledUsersByCourse :many
SELECT
    au.first_name,
    au.remaining_names,
    au.last_names,
    au.email
FROM attendee_course ac
JOIN attendee a ON ac.attendee_id = a.id
JOIN account_user au ON a.account_user_id = au.id
WHERE ac.course_id = $1;

-- name: DeleteAttendeeCourse :exec
DELETE FROM attendee_course
WHERE attendee_id = $1 AND course_id = $2;
