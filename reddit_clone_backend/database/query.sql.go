// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const checkForExistingUser = `-- name: CheckForExistingUser :one
SELECT EXISTS(SELECT 1 FROM reddit_clone.users WHERE email = ? OR user_name = ?)
`

type CheckForExistingUserParams struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

func (q *Queries) CheckForExistingUser(ctx context.Context, arg CheckForExistingUserParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkForExistingUser, arg.Email, arg.UserName)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createComment = `-- name: CreateComment :execresult

INSERT INTO reddit_clone.comments (
    comment, id_post, parent_id, id_user
) VALUES (
    ?, ?, ?, ?
)
`

type CreateCommentParams struct {
	Comment  string `json:"comment"`
	IDPost   int32  `json:"id_post"`
	ParentID *int32 `json:"parent_id"`
	IDUser   int32  `json:"id_user"`
}

// -------------------------------------------------- --
// Post and Comment
func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createComment,
		arg.Comment,
		arg.IDPost,
		arg.ParentID,
		arg.IDUser,
	)
}

const createPost = `-- name: CreatePost :execresult
INSERT INTO reddit_clone.posts (
    title, link, text, id_subreddit, id_user
) VALUES (
    ?, ?, ?, ?, ?
)
`

type CreatePostParams struct {
	Title       string  `json:"title"`
	Link        *string `json:"link"`
	Text        *string `json:"text"`
	IDSubreddit int32   `json:"id_subreddit"`
	IDUser      int32   `json:"id_user"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createPost,
		arg.Title,
		arg.Link,
		arg.Text,
		arg.IDSubreddit,
		arg.IDUser,
	)
}

const createSubreddit = `-- name: CreateSubreddit :execresult
INSERT INTO reddit_clone.subreddits (
    url, name, description, id_user
) VALUES (
    ?, ?, ?, ?
)
`

type CreateSubredditParams struct {
	Url         string  `json:"url"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	IDUser      *int32  `json:"id_user"`
}

func (q *Queries) CreateSubreddit(ctx context.Context, arg CreateSubredditParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createSubreddit,
		arg.Url,
		arg.Name,
		arg.Description,
		arg.IDUser,
	)
}

const createUser = `-- name: CreateUser :execresult

INSERT INTO reddit_clone.users (
    user_name, email, hashed_password
) VALUES (
    ?, ?, ?
)
`

type CreateUserParams struct {
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

// -------------------------------------------------- --
// Users
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.UserName, arg.Email, arg.HashedPassword)
}

const deletePost = `-- name: DeletePost :execresult
DELETE FROM reddit_clone.posts
WHERE id_post = ?
`

func (q *Queries) DeletePost(ctx context.Context, idPost int32) (sql.Result, error) {
	return q.db.ExecContext(ctx, deletePost, idPost)
}

const deleteUser = `-- name: DeleteUser :execresult
DELETE FROM reddit_clone.users
WHERE id_user = ?
`

func (q *Queries) DeleteUser(ctx context.Context, idUser int32) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteUser, idUser)
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id_user, user_name, email, hashed_password FROM reddit_clone.users
WHERE email = ? LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (RedditCloneUser, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i RedditCloneUser
	err := row.Scan(
		&i.IDUser,
		&i.UserName,
		&i.Email,
		&i.HashedPassword,
	)
	return i, err
}

const postByIdDetails = `-- name: PostByIdDetails :one

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
WHERE posts.id_post = ?
`

type PostByIdDetailsRow struct {
	IDPost    int32      `json:"id_post"`
	Title     string     `json:"title"`
	Link      *string    `json:"link"`
	Text      *string    `json:"text"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	UserName  string     `json:"user_name"`
	Score     int32      `json:"score"`
}

// -------------------------------------------------- --
// Posts
func (q *Queries) PostByIdDetails(ctx context.Context, idPost int32) (PostByIdDetailsRow, error) {
	row := q.db.QueryRowContext(ctx, postByIdDetails, idPost)
	var i PostByIdDetailsRow
	err := row.Scan(
		&i.IDPost,
		&i.Title,
		&i.Link,
		&i.Text,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserName,
		&i.Score,
	)
	return i, err
}

const postByIdListComments = `-- name: PostByIdListComments :many
SELECT comments.id_comment, comments.parent_id, comments.id_user, comments.comment
FROM reddit_clone.comments AS comments
WHERE comments.id_post = ?
`

type PostByIdListCommentsRow struct {
	IDComment int32  `json:"id_comment"`
	ParentID  *int32 `json:"parent_id"`
	IDUser    int32  `json:"id_user"`
	Comment   string `json:"comment"`
}

func (q *Queries) PostByIdListComments(ctx context.Context, idPost int32) ([]PostByIdListCommentsRow, error) {
	rows, err := q.db.QueryContext(ctx, postByIdListComments, idPost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PostByIdListCommentsRow
	for rows.Next() {
		var i PostByIdListCommentsRow
		if err := rows.Scan(
			&i.IDComment,
			&i.ParentID,
			&i.IDUser,
			&i.Comment,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const subredditByIdListPostsSortDate = `-- name: SubredditByIdListPostsSortDate :many
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
ORDER BY posts.created_at DESC
`

type SubredditByIdListPostsSortDateRow struct {
	IDPost    int32      `json:"id_post"`
	Title     string     `json:"title"`
	Link      *string    `json:"link"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	UserName  string     `json:"user_name"`
	Score     int32      `json:"score"`
}

func (q *Queries) SubredditByIdListPostsSortDate(ctx context.Context, idSubreddit int32) ([]SubredditByIdListPostsSortDateRow, error) {
	rows, err := q.db.QueryContext(ctx, subredditByIdListPostsSortDate, idSubreddit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SubredditByIdListPostsSortDateRow
	for rows.Next() {
		var i SubredditByIdListPostsSortDateRow
		if err := rows.Scan(
			&i.IDPost,
			&i.Title,
			&i.Link,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserName,
			&i.Score,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const subredditByIdListPostsSortScore = `-- name: SubredditByIdListPostsSortScore :many
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
ORDER BY score DESC
`

type SubredditByIdListPostsSortScoreRow struct {
	IDPost    int32      `json:"id_post"`
	Title     string     `json:"title"`
	Link      *string    `json:"link"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	UserName  string     `json:"user_name"`
	Score     int32      `json:"score"`
}

func (q *Queries) SubredditByIdListPostsSortScore(ctx context.Context, idSubreddit int32) ([]SubredditByIdListPostsSortScoreRow, error) {
	rows, err := q.db.QueryContext(ctx, subredditByIdListPostsSortScore, idSubreddit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SubredditByIdListPostsSortScoreRow
	for rows.Next() {
		var i SubredditByIdListPostsSortScoreRow
		if err := rows.Scan(
			&i.IDPost,
			&i.Title,
			&i.Link,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserName,
			&i.Score,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const subredditByUrlDetails = `-- name: SubredditByUrlDetails :one


SELECT subreddits.id_subreddit, subreddits.name, subreddits.description, users.user_name
FROM reddit_clone.subreddits AS subreddits
JOIN (reddit_clone.users AS users)
ON (users.id_user = subreddits.id_user)
WHERE subreddits.url = ?
`

type SubredditByUrlDetailsRow struct {
	IDSubreddit int32   `json:"id_subreddit"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	UserName    string  `json:"user_name"`
}

// Client:
// vis i realtid hvordan det lowercasede subreddit url kommer til at se ud imens du skriver display navnet
// -------------------------------------------------- --
// Subreddits
func (q *Queries) SubredditByUrlDetails(ctx context.Context, url string) (SubredditByUrlDetailsRow, error) {
	row := q.db.QueryRowContext(ctx, subredditByUrlDetails, url)
	var i SubredditByUrlDetailsRow
	err := row.Scan(
		&i.IDSubreddit,
		&i.Name,
		&i.Description,
		&i.UserName,
	)
	return i, err
}

const voteForPostOrComment = `-- name: VoteForPostOrComment :execresult
INSERT INTO reddit_clone.scores (
    score, id_post, id_comment, id_user
) VALUES (
    ?, ?, ?, ?
)
`

type VoteForPostOrCommentParams struct {
	Score     bool   `json:"score"`
	IDPost    *int32 `json:"id_post"`
	IDComment *int32 `json:"id_comment"`
	IDUser    int32  `json:"id_user"`
}

func (q *Queries) VoteForPostOrComment(ctx context.Context, arg VoteForPostOrCommentParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, voteForPostOrComment,
		arg.Score,
		arg.IDPost,
		arg.IDComment,
		arg.IDUser,
	)
}
