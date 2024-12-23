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

-- name: SubredditByIdListPostsSortNewestStart :many
SELECT
posts.id_post,
posts.title,
posts.link,
posts.created_at,
posts.updated_at,
users.user_name,
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = TRUE AND scores.id_post = posts.id_post) -
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = FALSE AND scores.id_post = posts.id_post)
AS score,
(SELECT COUNT(*) FROM reddit_clone.comments WHERE comments.id_post = posts.id_post)
as comments_count
FROM reddit_clone.posts AS posts
JOIN (reddit_clone.users AS users)
ON (users.id_user = posts.id_user)
WHERE posts.id_subreddit = ?
ORDER BY posts.id_post DESC
LIMIT 30;

-- name: SubredditByIdListPostsSortNewestKeySetPaginated :many
SELECT
posts.id_post,
posts.title,
posts.link,
posts.created_at,
posts.updated_at,
users.user_name,
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = TRUE AND scores.id_post = posts.id_post) -
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = FALSE AND scores.id_post = posts.id_post)
AS score,
(SELECT COUNT(*) FROM reddit_clone.comments WHERE comments.id_post = posts.id_post)
as comments_count
FROM reddit_clone.posts AS posts
JOIN (reddit_clone.users AS users)
ON (users.id_user = posts.id_user)
WHERE posts.id_subreddit = ? AND posts.id_post < ? 
ORDER BY posts.id_post DESC
LIMIT 30;

-- name: SubredditByIdListPostsSortScoreHighest :many
SELECT
posts.id_post,
posts.title,
posts.link,
posts.created_at,
posts.updated_at,
users.user_name,
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = TRUE AND scores.id_post = posts.id_post) -
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = FALSE AND scores.id_post = posts.id_post)
AS score,
(SELECT COUNT(*) FROM reddit_clone.comments WHERE comments.id_post = posts.id_post)
as comments_count
FROM reddit_clone.posts AS posts
JOIN (reddit_clone.users AS users)
ON (users.id_user = posts.id_user)
WHERE posts.id_subreddit = ?
ORDER BY score DESC
LIMIT 30 OFFSET ?;

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
AS score,
subreddits.name
FROM reddit_clone.posts AS posts
JOIN (reddit_clone.users AS users, reddit_clone.subreddits AS subreddits)
ON (users.id_user = posts.id_user AND subreddits.id_subreddit = posts.id_subreddit)
WHERE posts.id_post = ?;

-- name: PostByIdListComments :many
SELECT 
comments.id_comment,
comments.parent_id,
comments.created_at,
comments.updated_at,
comments.comment,
users.user_name,
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = TRUE AND scores.id_comment = comments.id_comment) -
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = FALSE AND scores.id_comment = comments.id_comment)
AS score
FROM reddit_clone.comments AS comments
JOIN (reddit_clone.users AS users)
ON (users.id_user = comments.id_user)
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

-- name: PostsFullTextSearch :many
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
WHERE MATCH (posts.title, posts.text)
AGAINST (? IN NATURAL LANGUAGE MODE);

                                                      --
-- -------------------------------------------------- --
-- Comment

-- name: GetCommentById :one
SELECT 
comments.id_comment,
comments.parent_id,
comments.created_at,
comments.updated_at,
comments.comment,
users.user_name,
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = TRUE AND scores.id_comment = comments.id_comment) -
(SELECT COUNT(*) FROM reddit_clone.scores WHERE scores.score = FALSE AND scores.id_comment = comments.id_comment)
AS score
FROM reddit_clone.comments AS comments
JOIN (reddit_clone.users AS users)
ON (users.id_user = comments.id_user)
WHERE comments.id_comment = ?;

-- name: CreateCommentOnPost :execresult
INSERT INTO reddit_clone.comments (
    comment, id_post, id_user, is_toplevel
) VALUES (
    ?, ?, ?, true
);

-- name: CreateCommentOnParent :execresult
INSERT INTO reddit_clone.comments (
    comment, id_post, parent_id, id_user, is_toplevel
) VALUES (
    ?, ?, ?, ?, false
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

                                                      --
-- -------------------------------------------------- --
-- Scores

-- name: VoteOnPost :execresult
INSERT INTO reddit_clone.scores (
    score, id_post, id_user
) VALUES (
    ?, ?, ?
);

-- name: VoteOnComment :execresult
INSERT INTO reddit_clone.scores (
    score, id_comment, id_user
) VALUES (
    ?, ?, ?
);