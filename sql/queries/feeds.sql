-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetFeeds :many
SELECT feeds.id, feeds.created_at, feeds.updated_at, feeds.name, feeds.url, users.name as username
FROM feeds JOIN users ON feeds.user_id = users.id;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at=NOW(), fetched_at=NOW()
WHERE id=$1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY fetched_at ASC NULLS FIRST
LIMIT 1;