-- name: CreateStudentGroup :exec
INSERT INTO student_group (
  name, priority, start_day, end_day
) VALUES (
  $1, $2, $3, $4
);
