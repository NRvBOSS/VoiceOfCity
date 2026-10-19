
-- name: CreateIssue :one
INSERT INTO issues (
    title,
    description,
    category,
    latitude,
    longitude
)
VALUES ($1, $2, $3, $4, $5)
RETURNING
    id,
    title,
    description,
    category,
    latitude,
    longitude,
    status,
    created_at;


-- name: GetIssue :one
SELECT
    id,
    title,
    description,
    category,
    latitude,
    longitude,
    status,
    created_at
FROM issues
WHERE id = $1;


-- name: ListIssues :many
SELECT
    id,
    title,
    description,
    category,
    latitude,
    longitude,
    status,
    created_at
FROM issues
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
