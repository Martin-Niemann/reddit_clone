package main

import (
	"time"
)

type Comment struct {
	Id       int32  `json:"id"`
	ParentId *int32 `json:"parent_id"`
	UserId   int32  `json:"user_id"`
	Comment  string `json:"comment"`
}

type PostCard struct {
	Id        int32      `json:"id_post"`
	Title     string     `json:"title"`
	Link      *string    `json:"link"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	UserName  string     `json:"username"`
	Score     int32      `json:"score"`
}

type PostFull struct {
	PostCard
	Text     *string   `json:"text"`
	Comments []Comment `json:"comments"`
}

type Subreddit struct {
	Id          int32      `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Moderator   string     `json:"moderator"`
	Posts       []PostCard `json:"posts"`
}

type ServiceErrorType int

const (
	NoResult ServiceErrorType = iota
	UnexpectedError
	NoError
)

type ServiceError struct {
	Type ServiceErrorType
	Text string
}
