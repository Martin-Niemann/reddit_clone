-- -------------------------------------------------- --
-- Subreddit

-- name: SubredditByUrlDetails :one
SELECT subreddits.id_subreddit, subreddits.name, subreddits.description, users.user_name
FROM reddit_clone.subreddits AS subreddits
JOIN (reddit_clone.users AS users)
ON (users.id_user = subreddits.id_user)
WHERE subreddits.url = ?;

-- name: SubredditByIdListPostsSortDate :many
SELECT posts.title, posts.link, posts.created_at, posts.updated_at, users.user_name, scores.score FROM reddit_clone.posts AS posts
JOIN (reddit_clone.scores AS scores, reddit_clone.users AS users)
ON (scores.id_post = posts.id_post, users.id_user = posts.id_user)
WHERE posts.id_subreddit = ?
ORDER BY posts.created_at;

-- name: SubredditByIdListPostsSortScore :many
SELECT posts.title, posts.link, posts.created_at, posts.updated_at, users.user_name, scores.score FROM reddit_clone.posts AS posts
JOIN (reddit_clone.scores AS scores, reddit_clone.users AS users)
ON (scores.id_post = posts.id_post, users.id_user = posts.id_user)
WHERE posts.id_subreddit = ?
ORDER BY scores.score;
                                                      --
-- -------------------------------------------------- --
-- Users

-- name: CreatePost :execresult
INSERT INTO reddit_clone.posts (
    title, link, text, id_subreddit, id_user
) VALUES (
    ?, ?, ?, ?, ?
);



-- name: DeletePost :exec
DELETE FROM reddit_clone.posts
WHERE id_post = ?;


-- -----------------------------------------------------
-- Posts

-- name: GetPost :one
SELECT * FROM reddit_clone.posts
WHERE id_post = ? LIMIT 1;
