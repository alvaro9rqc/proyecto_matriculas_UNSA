-- name: CreateAttendeeGroup :exec
INSERT INTO attendee_group (
  name, priority
) VALUES (
  $1, $2
);
