-- name: CreateSpeaker :exec
INSERT INTO speaker (
    account_id
) VALUES (
    $1
);

-- name: ListSpeakers :many
SELECT
    s.id,
    a.email,
    a.name,
    a.surname,
    a.avatar_url
FROM speaker s
JOIN account a ON s.account_id = a.id
ORDER BY a.name, a.surname
LIMIT $1
OFFSET $2;

-- name: FullListSpeakers :many
SELECT
    s.id,
    a.email,
    a.name,
    a.surname,
    a.avatar_url
FROM speaker s
JOIN account a ON s.account_id = a.id
ORDER BY a.name, a.surname;

-- name: ListMajorBySpeaker :many
SELECT DiSTINCT
    m.id AS major_id,
    m.name AS major_name
FROM speaker s
JOIN section_speaker ss ON s.id = ss.speaker_id
JOIN section sec ON ss.section_id = sec.id
JOIN course c ON sec.course_id = c.id
JOIN major m ON c.major_id = m.id
WHERE s.account_id = $1;
