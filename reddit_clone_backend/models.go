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

type Token struct {
	Token string `json:"token" validate:"jwt"`
}

type SignUpDetails struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=14,max=72"`
	UserName string `json:"username" validate:"required,min=1,max=100"`
}

type UserId struct {
	Id int32 `json:"id"`
}

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}
