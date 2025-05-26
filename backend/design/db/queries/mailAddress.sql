-- name: ListAllMailAddresses :many
SELECT
    *
FROM
    mailAddress
ORDER BY
    updatedAt DESC;

-- name: ListMailAddresses :many
SELECT
    *
FROM
    mailAddress
ORDER BY
    updatedAt DESC
LIMIT
    ?
OFFSET
    ?;

-- name: GetMailAddress :one
SELECT
    *
FROM
    mailAddress
WHERE
    id = ?;

-- name: CreateMailAddress :one
INSERT INTO
    mailAddress (id, address)
VALUES
    (?, ?) RETURNING *;

-- name: DeleteMailAddress :exec
DELETE FROM mailAddress
WHERE
    id = ?;

-- name: UpdateMailAddress :exec
UPDATE mailAddress
SET
    address = ?,
    active = ?,
    updatedAt = CURRENT_TIMESTAMP
WHERE
    id = ?;
