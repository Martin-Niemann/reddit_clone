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

type GetSubredditResult struct {
	Details database.SubredditByUrlDetailsRow            `json:"details"`
	Posts   []database.SubredditByIdListPostsSortDateRow `json:"posts"`
}

func (s *Service) getSubreddit(ctx context.Context, url string) (*GetSubredditResult, error) {
	subredditDetails, err := s.queries.SubredditByUrlDetails(ctx, url)
	if err != nil {
		return nil, err
	}

	subredditPosts, err := s.queries.SubredditByIdListPostsSortDate(ctx, subredditDetails.IDSubreddit)
	if err != nil {
		return nil, err
	}

	return &GetSubredditResult{Details: subredditDetails, Posts: subredditPosts}, nil
}
