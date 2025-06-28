-- name: CreateProcess :exec
INSERT INTO process (name, start_day, end_day) VALUES ($1, $2, $3);

-- name: ListAllProcess :many
SELECT id, name, start_day, end_day
FROM process
ORDER BY start_day;

-- name: DeleteProcess :exec
DELETE FROM process WHERE id = $1;
