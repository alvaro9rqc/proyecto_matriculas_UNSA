-- name: CreateProcess :exec
INSERT INTO process (name, start_day, end_day, institution_id) VALUES ($1, $2, $3, $4);

-- name: ListAllProcess :many
SELECT *
FROM process
ORDER BY start_day;

-- name: DeleteProcess :exec
DELETE FROM process WHERE id = $1;
