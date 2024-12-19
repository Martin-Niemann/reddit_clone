// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"time"
)

type RedditCloneComment struct {
	IDComment  int32      `json:"id_comment"`
	Comment    string     `json:"comment"`
	IsToplevel bool       `json:"is_toplevel"`
	IDPost     int32      `json:"id_post"`
	ParentID   *int32     `json:"parent_id"`
	IDUser     int32      `json:"id_user"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type RedditClonePost struct {
	IDPost      int32      `json:"id_post"`
	Title       string     `json:"title"`
	Link        *string    `json:"link"`
	Text        *string    `json:"text"`
	IDSubreddit int32      `json:"id_subreddit"`
	IDUser      int32      `json:"id_user"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type RedditCloneScore struct {
	IDScore   int32      `json:"id_score"`
	Score     bool       `json:"score"`
	IDPost    *int32     `json:"id_post"`
	IDComment *int32     `json:"id_comment"`
	IDUser    int32      `json:"id_user"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type RedditCloneSubreddit struct {
	IDSubreddit int32   `json:"id_subreddit"`
	Url         string  `json:"url"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	IDUser      *int32  `json:"id_user"`
}

type RedditCloneUser struct {
	IDUser         int32  `json:"id_user"`
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}
