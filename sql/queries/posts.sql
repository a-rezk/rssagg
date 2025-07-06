-- name: CreatePost :execresult
INSERT IGNORE INTO posts (id, created_at, updated_at, published_at, title, description, url, feed_id)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetPostByID :one
SELECT * FROM posts
WHERE id = ?;

-- name: GetPostsForUser :many
SELECT posts.* FROM posts 
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id 
WHERE feed_follows.user_id = ? 
ORDER by posts.published_at DESC 
LIMIT ?;