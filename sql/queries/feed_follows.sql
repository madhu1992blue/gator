-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS(
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES($1, $2, $3, (SELECT id FROM users WHERE users.name=sqlc.arg(username)), (SELECT id FROM feeds WHERE feeds.url=sqlc.arg(feed_url))) RETURNING *
)
SELECT inserted_feed_follow.*, feeds.name AS feed_name, users.name AS user_name
FROM 
    inserted_feed_follow 
    INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
    INNER JOIN users ON inserted_feed_follow.user_id = users.id;


-- name: GetFeedFollowsForUser :many
SELECT feeds.name as feed_name, users.name as user_name
FROM feed_follows 
    INNER JOIN feeds ON feed_follows.feed_id = feeds.id
    INNER JOIN users ON feed_follows.user_id = users.id
    WHERE users.name=$1;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows
WHERE user_id=(SELECT id FROM users WHERE users.name=sqlc.arg(username))
      AND feed_id=(SELECT id from feeds WHERE feeds.url=sqlc.arg(feed_url));