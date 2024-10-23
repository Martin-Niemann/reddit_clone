-- name: GetPost :one
SELECT * FROM reddit_clone.posts
WHERE id_post = ? LIMIT 1;

-- name: ListPosts :many
SELECT * FROM reddit_clone.posts
ORDER BY created_at;

-- name: CreatePost :execresult
INSERT INTO reddit_clone.posts (
    title, link, text, id_subreddit, id_user
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: DeletePost :exec
DELETE FROM reddit_clone.posts
WHERE id_post = ?;
