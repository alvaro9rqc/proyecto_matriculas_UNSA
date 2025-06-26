-- name: ListMajorByAdmin :one
SELECT 
    m.id AS major_id,
    m.name AS major_name
FROM major m 
JOIN administrative a ON m.id = a.major_id
WHERE a.account_id = $1;
