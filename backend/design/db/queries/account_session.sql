-- name: CreateAccountSession :one
INSERT INTO account_session (
    token,
    expiration_date,
    user_agent,
    ip_address,
    account_id
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, token, expiration_date, user_agent, ip_address, account_id;
