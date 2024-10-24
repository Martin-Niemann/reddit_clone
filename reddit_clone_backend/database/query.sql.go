// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :execresult
INSERT INTO reddit_clone.post (
    title, link, text, id_subreddit, id_user
) VALUES (
    ?, ?, ?, ?, ?
)
`

type CreatePostParams struct {
	Title       string
	Link        sql.NullString
	Text        sql.NullString
	IDSubreddit int32
	IDUser      int32
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

const deletePost = `-- name: DeletePost :exec
DELETE FROM reddit_clone.post
WHERE id_post = ?
`

func (q *Queries) DeletePost(ctx context.Context, idPost int32) error {
	_, err := q.db.ExecContext(ctx, deletePost, idPost)
	return err
}

const getPost = `-- name: GetPost :one
SELECT id_post, title, link, text, created_date, edited_date, id_subreddit, id_user FROM reddit_clone.post
WHERE id_post = ? LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, idPost int32) (RedditClonePost, error) {
	row := q.db.QueryRowContext(ctx, getPost, idPost)
	var i RedditClonePost
	err := row.Scan(
		&i.IDPost,
		&i.Title,
		&i.Link,
		&i.Text,
		&i.CreatedDate,
		&i.EditedDate,
		&i.IDSubreddit,
		&i.IDUser,
	)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT id_post, title, link, text, created_date, edited_date, id_subreddit, id_user FROM reddit_clone.post
ORDER BY created_date
`

func (q *Queries) ListPosts(ctx context.Context) ([]RedditClonePost, error) {
	rows, err := q.db.QueryContext(ctx, listPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RedditClonePost
	for rows.Next() {
		var i RedditClonePost
		if err := rows.Scan(
			&i.IDPost,
			&i.Title,
			&i.Link,
			&i.Text,
			&i.CreatedDate,
			&i.EditedDate,
			&i.IDSubreddit,
			&i.IDUser,
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
