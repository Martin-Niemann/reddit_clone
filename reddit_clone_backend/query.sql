-- name: GetPost :one
SELECT * FROM reddit_clone.post
WHERE id_post = ? LIMIT 1;

-- name: ListPosts :many
SELECT * FROM reddit_clone.post
ORDER BY created_date;

-- name: CreatePost :execresult
INSERT INTO reddit_clone.post (
    title, link, text, id_subreddit, id_user
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: DeletePost :exec
DELETE FROM reddit_clone.post
WHERE id_post = ?;