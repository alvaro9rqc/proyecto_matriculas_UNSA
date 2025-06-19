-- name: GetAccountByEmail :one
SELECT * FROM account
WHERE email = $1;
