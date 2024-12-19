package main

import (
	"time"
)

type Comment struct {
	Id        int32      `json:"id"`
	ParentId  *int32     `json:"parent_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	UserName  string     `json:"username"`
	Text      string     `json:"text"`
	Score     int32      `json:"score"`
}

type CommentDetails struct {
	PostId   int32  `json:"post_id" validate:"required"`
	ParentId int32  `json:"parent_id"`
	UserId   int32  `json:"user_id" validate:"required"`
	Text     string `json:"text" validate:"required"`
}

type ScoreDetails struct {
	Score     int32 `json:"score"`
	PostId    int32 `json:"post_id"`
	CommentId int32 `json:"comment_id"`
	UserId    int32 `json:"user_id"`
}

type PostCard struct {
	Id            int32      `json:"id"`
	Title         string     `json:"title"`
	Link          *string    `json:"link"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	UserName      string     `json:"username"`
	Score         int32      `json:"score"`
	CommentsCount int64      `json:"comments_count"`
}

type PostFull struct {
	PostCard
	Text          *string   `json:"text"`
	Comments      []Comment `json:"comments"`
	CommunityName string    `json:"community_name"`
}

type Subreddit struct {
	Id          int32      `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Moderator   string     `json:"moderator"`
	Posts       []PostCard `json:"posts"`
}

type SubredditDetails struct {
	Url         string `json:"url"`
	DisplayName string `json:"displayname"`
	Description string `json:"description"`
	UserId      int32
}

type PostDetails struct {
	Title       string  `json:"title"`
	Link        *string `json:"link,omitempty"`
	Text        *string `json:"text,omitempty"`
	CommunityId int32   `json:"community_id"`
	UserId      int32
}

type PostDetailsOnlyTitle struct {
	Title       string `json:"title"`
	CommunityId int32  `json:"community_id"`
	UserId      int32
}

type ServiceErrorType int

const (
	NoResult ServiceErrorType = iota
	InvalidInput
	InvalidArgument
	UnexpectedError
	UserAlreadyExists
	UserDoesntExist
	InvalidEmailAndOrPassword
	MissingAuthCookie
	InvalidOrExpiredToken
	NoError
)

type ServiceError struct {
	Type             ServiceErrorType
	Text             string
	ValidationErrors []ValidationError
}

type LoginDetails struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=14,max=72"`
}

type SignUpDetails struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
	UserName string `json:"username" validate:"required,min=4,max=100"`
}

type UserId struct {
	Id int32 `json:"id"`
}

type CommentId struct {
	Id int64 `json:"id"`
}

type PostId struct {
	Id int64 `json:"post_id"`
}

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type ReceivedToken struct {
	Token string `json:"token" validate:"jwt"`
}

type TokenCookie struct {
	Token   string `json:"token"`
	Expires int64  `json:"expires"`
}
