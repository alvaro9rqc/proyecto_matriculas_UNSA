-- name: CreateStudentGroup :exec
INSERT INTO student_group (
  name, priority
) VALUES (
  $1, $2
);
