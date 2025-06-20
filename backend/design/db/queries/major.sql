CREATE TABLE major (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) UNIQUE NOT NULL
);

-- name: CreateMajor :exec
INSERT INTO major (name)
VALUES ($1);

-- name: UpdateMajor :exec
UPDATE major
SET
    name = $1
WHERE id = $2;

-- name: DeleteMajor :exec
DELETE FROM major
WHERE id = $1;

-- name: ListMajors :many
SELECT
    id,
    name
FROM major
ORDER BY name;
