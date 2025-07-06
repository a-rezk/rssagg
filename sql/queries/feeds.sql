-- name: CreateFeed :execresult
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetFeedByID :one
SELECT * FROM feeds
WHERE id = ?;

-- name: DeleteFeedByID :exec
DELETE FROM feeds WHERE id = ? AND user_id = ?;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC LIMIT ?;

-- name: MarkFeedAsFetched :exec
UPDATE feeds 
SET 
    last_fetched_at = NOW(),
    updated_at = NOW()
WHERE id = ?;

