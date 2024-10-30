package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reddit_clone_backend/database"
	"time"
)

type Service struct {
	db      *sql.DB
	queries *database.Queries
}

type MysqlConnectionVariables struct {
	MYSQL_HOST string
	MYSQL_PORT string
	MYSQL_USER string
	MYSQL_PASS string
}

func (s *Service) closeConnection() {
	s.db.Close()
}

func getServiceErrorConst(err error) ServiceError {
	if err == sql.ErrNoRows {
		return ServiceError{Type: NoResult}
	}

	log.Println("A new error! ", err.Error())
	return ServiceError{Type: UnexpectedError, Text: err.Error()}
}

func NewService(mysql MysqlConnectionVariables) Service {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/reddit_clone?parseTime=true", mysql.MYSQL_USER, mysql.MYSQL_PASS, mysql.MYSQL_HOST, mysql.MYSQL_PORT)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err.Error())
	}
	log.Println("Database connection succesful.")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	queries := database.New(db)
	return Service{db: db, queries: queries}
}

func (s *Service) getSubredditSortedByDate(ctx context.Context, url string) (*Subreddit, ServiceError) {
	subredditDetails, err := s.queries.SubredditByUrlDetails(ctx, url)
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	subredditPosts, err := s.queries.SubredditByIdListPostsSortDate(ctx, subredditDetails.IDSubreddit)
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	posts := []PostCard{}

	for _, post := range subredditPosts {
		posts = append(
			posts,
			PostCard{
				Id:        post.IDPost,
				Title:     post.Title,
				Link:      post.Link,
				CreatedAt: post.CreatedAt,
				UpdatedAt: post.UpdatedAt,
				UserName:  post.UserName,
				Score:     post.Score})
	}

	data := Subreddit{
		Id:          subredditDetails.IDSubreddit,
		Name:        subredditDetails.Name,
		Description: subredditDetails.Description,
		Moderator:   subredditDetails.UserName,
		Posts:       posts,
	}

	return &data, ServiceError{Type: NoError}
}

func (s *Service) getSubredditSortedByScore(ctx context.Context, url string) (*Subreddit, ServiceError) {
	subredditDetails, err := s.queries.SubredditByUrlDetails(ctx, url)
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	subredditPosts, err := s.queries.SubredditByIdListPostsSortScore(ctx, subredditDetails.IDSubreddit)
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	posts := []PostCard{}

	for _, post := range subredditPosts {
		posts = append(
			posts,
			PostCard{
				Id:        post.IDPost,
				Title:     post.Title,
				Link:      post.Link,
				CreatedAt: post.CreatedAt,
				UpdatedAt: post.UpdatedAt,
				UserName:  post.UserName,
				Score:     post.Score})
	}

	data := Subreddit{
		Id:          subredditDetails.IDSubreddit,
		Name:        subredditDetails.Name,
		Description: subredditDetails.Description,
		Moderator:   subredditDetails.UserName,
		Posts:       posts,
	}

	return &data, ServiceError{Type: NoError}
}

func (s *Service) getPost(ctx context.Context, id int32) (*PostFull, ServiceError) {
	postDetails, err := s.queries.PostByIdDetails(ctx, id)
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	postComments, err := s.queries.PostByIdListComments(ctx, id)
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	comments := []Comment{}

	for _, comment := range postComments {
		comments = append(
			comments,
			Comment{
				Id:       comment.IDComment,
				ParentId: comment.ParentID,
				UserId:   comment.IDUser,
				Comment:  comment.Comment,
			})
	}

	data := PostFull{
		PostCard: PostCard{
			Id:        postDetails.IDPost,
			Title:     postDetails.Title,
			Link:      postDetails.Link,
			CreatedAt: postDetails.CreatedAt,
			UpdatedAt: postDetails.UpdatedAt,
			UserName:  postDetails.UserName,
			Score:     postDetails.Score,
		},
		Text:     postDetails.Text,
		Comments: comments,
	}

	return &data, ServiceError{Type: NoError}
}
