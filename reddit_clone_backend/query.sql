-- Client:
-- vis i realtid hvordan det lowercasede subreddit url kommer til at se ud imens du skriver display navnet

-- -------------------------------------------------- --
-- Subreddits

-- name: SubredditByUrlDetails :one
SELECT subreddits.id_subreddit, subreddits.name, subreddits.description, users.user_name
FROM reddit_clone.subreddits AS subreddits
JOIN (reddit_clone.users AS users)
ON (users.id_user = subreddits.id_user)
WHERE subreddits.url = ?;

-- name: SubredditByIdListPostsSortDate :many
SELECT
posts.id_post,
posts.title,
posts.link,
posts.created_at,
posts.updated_at,
users.user_name,
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = TRUE AND scores.id_post = posts.id_post) -
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = FALSE AND scores.id_post = posts.id_post)
AS score
FROM reddit_clone.posts AS posts
JOIN (reddit_clone.users AS users)
ON (users.id_user = posts.id_user)
WHERE posts.id_subreddit = ?
ORDER BY posts.created_at DESC;

-- name: SubredditByIdListPostsSortScore :many
SELECT
posts.id_post,
posts.title,
posts.link,
posts.created_at,
posts.updated_at,
users.user_name,
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = TRUE AND scores.id_post = posts.id_post) -
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = FALSE AND scores.id_post = posts.id_post)
AS score
FROM reddit_clone.posts AS posts
JOIN (reddit_clone.users AS users)
ON (users.id_user = posts.id_user)
WHERE posts.id_subreddit = ?
ORDER BY score DESC;

-- name: CreateSubreddit :execresult
INSERT INTO reddit_clone.subreddits (
    url, name, description, id_user
) VALUES (
    ?, ?, ?, ?
);
                                                      --
-- -------------------------------------------------- --
-- Posts

-- name: PostByIdDetails :one
SELECT
posts.id_post,
posts.title,
posts.link,
posts.text,
posts.created_at,
posts.updated_at,
users.user_name,
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = TRUE AND scores.id_post = posts.id_post) -
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = FALSE AND scores.id_post = posts.id_post)
AS score
FROM reddit_clone.posts AS posts
JOIN (reddit_clone.users AS users)
ON (users.id_user = posts.id_user)
WHERE posts.id_post = ?;

-- name: PostByIdListComments :many
SELECT comments.id_comment, comments.parent_id, comments.id_user, comments.comment
FROM reddit_clone.comments AS comments
WHERE comments.id_post = ?;

-- name: CreatePost :execresult
INSERT INTO reddit_clone.posts (
    title, link, text, id_subreddit, id_user
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: DeletePost :execresult
DELETE FROM reddit_clone.posts
WHERE id_post = ?;
                                                      --
-- -------------------------------------------------- --
-- Post and Comment

-- name: CreateComment :execresult
INSERT INTO reddit_clone.comments (
    comment, id_post, parent_id, id_user
) VALUES (
    ?, ?, ?, ?
);

-- name: VoteForPostOrComment :execresult
INSERT INTO reddit_clone.scores (
    score, id_post, id_comment, id_user
) VALUES (
    ?, ?, ?, ?
);
                                                      --
-- -------------------------------------------------- --
-- Users

-- name: CreateUser :execresult
INSERT INTO reddit_clone.users (
    user_name, email, hashed_password
) VALUES (
    ?, ?, ?
);

-- name: CheckForExistingUser :one
SELECT EXISTS(SELECT 1 FROM reddit_clone.users WHERE email = ? OR user_name = ?);

-- name: GetUserByEmail :one
SELECT * FROM reddit_clone.users
WHERE email = ? LIMIT 1;

-- name: DeleteUser :execresult
DELETE FROM reddit_clone.users
WHERE id_user = ?;
