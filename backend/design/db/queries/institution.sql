CREATE TABLE institution (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) UNIQUE NOT NULL,
    logo_url VARCHAR(255)
);

-- name: CreateInstitution :exec
INSERT INTO institution (name, logo_url) VALUES ($1, $2);

-- name: ListAllInstitutions :many
SELECT *
FROM institution
ORDER BY name;


-- name: ListInstitutionsByAccountID :many
SELECT i.*
FROM institution i
JOIN process p ON i.id = p.institution_id
JOIN student_process sp ON p.id = sp.process_id
JOIN student s ON sp.student_id = s.id
JOIN account a ON s.account_id = a.id
WHERE a.id = $1
ORDER BY i.name;
