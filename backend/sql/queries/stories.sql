-- name: CreateStory :one
INSERT INTO stories(title)
VALUES (sqlc.arg(title))
RETURNING *;




