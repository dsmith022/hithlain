-- name: CreateNode :one

INSERT INTO nodes(
title,
body
)
VALUES(
    sqlc.arg(title),
    sqlc.arg(body)
)
RETURNING *;


-- name: CreateNodeStoryRelationship :one

INSERT INTO story_nodes(
    node_id,
    story_id
)
VALUES(
    sqlc.arg(node_id),
    sqlc.arg(story_id)
)
RETURNING *;