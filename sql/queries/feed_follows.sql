-- name: CreateFeedFollow :exec
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (?, ?, ?, ?, ?);

-- name: GetFeedFollow :one
SELECT * FROM feed_follows
WHERE id = ?;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows
WHERE user_id = ?;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE user_id = ? AND id = ?;