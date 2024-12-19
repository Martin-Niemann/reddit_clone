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

func (s *Service) getSubredditSortedByNewestStart(ctx context.Context, url string) (*Subreddit, ServiceError) {
	subredditDetails, err := s.queries.SubredditByUrlDetails(ctx, url)
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	subredditPosts, err := s.queries.SubredditByIdListPostsSortNewestStart(ctx, subredditDetails.IDSubreddit)
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	posts := []PostCard{}

	for _, post := range subredditPosts {
		posts = append(
			posts,
			PostCard{
				Id:            post.IDPost,
				Title:         post.Title,
				Link:          post.Link,
				CreatedAt:     post.CreatedAt,
				UpdatedAt:     post.UpdatedAt,
				UserName:      post.UserName,
				Score:         post.Score,
				CommentsCount: post.CommentsCount})
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

func (s *Service) getSubredditSortedByNewestKeysetPaginated(ctx context.Context, url string, postId int32) (*Subreddit, ServiceError) {
	subredditDetails, err := s.queries.SubredditByUrlDetails(ctx, url)
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	subredditPosts, err := s.queries.SubredditByIdListPostsSortNewestKeySetPaginated(ctx, database.SubredditByIdListPostsSortNewestKeySetPaginatedParams{IDSubreddit: subredditDetails.IDSubreddit, IDPost: postId})
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	posts := []PostCard{}

	for _, post := range subredditPosts {
		posts = append(
			posts,
			PostCard{
				Id:            post.IDPost,
				Title:         post.Title,
				Link:          post.Link,
				CreatedAt:     post.CreatedAt,
				UpdatedAt:     post.UpdatedAt,
				UserName:      post.UserName,
				Score:         post.Score,
				CommentsCount: post.CommentsCount})
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

func (s *Service) getSubredditSortedByScoreHighestOffsetPaginated(ctx context.Context, url string, offset int32) (*Subreddit, ServiceError) {
	subredditDetails, err := s.queries.SubredditByUrlDetails(ctx, url)
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	subredditPosts, err := s.queries.SubredditByIdListPostsSortScoreHighest(ctx, database.SubredditByIdListPostsSortScoreHighestParams{IDSubreddit: subredditDetails.IDSubreddit, Offset: offset})
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	posts := []PostCard{}

	for _, post := range subredditPosts {
		posts = append(
			posts,
			PostCard{
				Id:            post.IDPost,
				Title:         post.Title,
				Link:          post.Link,
				CreatedAt:     post.CreatedAt,
				UpdatedAt:     post.UpdatedAt,
				UserName:      post.UserName,
				Score:         post.Score,
				CommentsCount: post.CommentsCount})
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
				Id:        comment.IDComment,
				ParentId:  comment.ParentID,
				CreatedAt: comment.CreatedAt,
				UpdatedAt: comment.UpdatedAt,
				UserName:  comment.UserName,
				Text:      comment.Comment,
				Score:     comment.Score,
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
		CommunityName: postDetails.Name,
		Text:          postDetails.Text,
		Comments:      comments,
	}

	return &data, ServiceError{Type: NoError}
}

func (s *Service) getComment(ctx context.Context, id int32) (*Comment, ServiceError) {
	dbComment, err := s.queries.GetCommentById(ctx, id)
	if err != nil {
		return nil, getServiceErrorConst(err)
	}

	comment := Comment{
		Id:        dbComment.IDComment,
		ParentId:  dbComment.ParentID,
		CreatedAt: dbComment.CreatedAt,
		UpdatedAt: dbComment.UpdatedAt,
		UserName:  dbComment.UserName,
		Text:      dbComment.Comment,
		Score:     dbComment.Score,
	}

	return &comment, ServiceError{Type: NoError}
}

func (s *Service) addComment(ctx context.Context, commentDetails *CommentDetails) (int64, ServiceError) {
	var result sql.Result
	var err error

	if commentDetails.ParentId != 0 {
		result, err = s.queries.CreateCommentOnParent(ctx, database.CreateCommentOnParentParams{Comment: commentDetails.Text, IDPost: commentDetails.PostId, ParentID: &commentDetails.ParentId, IDUser: commentDetails.UserId})
		if err != nil {
			return 0, getServiceErrorConst(err)
		}
	} else {
		result, err = s.queries.CreateCommentOnPost(ctx, database.CreateCommentOnPostParams{Comment: commentDetails.Text, IDPost: commentDetails.PostId, IDUser: commentDetails.UserId})
		if err != nil {
			return 0, getServiceErrorConst(err)
		}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, getServiceErrorConst(err)
	}

	return insertId, ServiceError{Type: NoError}
}

func (s *Service) voteOnComment(ctx context.Context, scoreDetails *ScoreDetails) (int64, ServiceError) {
	var result sql.Result
	var err error

	switch scoreDetails.Score {
	case -1:
		result, err = s.queries.VoteOnComment(ctx, database.VoteOnCommentParams{Score: false, IDComment: &scoreDetails.CommentId, IDUser: scoreDetails.UserId})
		if err != nil {
			return 0, getServiceErrorConst(err)
		}
	case 1:
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, getServiceErrorConst(err)
	}

	return insertId, ServiceError{Type: NoError}
}

func (s *Service) addCommunity(ctx context.Context, subredditDetails *SubredditDetails) (int64, ServiceError) {
	var result sql.Result
	var err error

	result, err = s.queries.CreateSubreddit(ctx, database.CreateSubredditParams{Url: subredditDetails.Url, Name: subredditDetails.DisplayName, Description: &subredditDetails.Description, IDUser: &subredditDetails.UserId})
	if err != nil {
		return 0, getServiceErrorConst(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, getServiceErrorConst(err)
	}

	return insertId, ServiceError{Type: NoError}
}

func (s *Service) addPost(ctx context.Context, postDetails *PostDetails) (int64, ServiceError) {
	var result sql.Result
	var err error

	result, err = s.queries.CreatePost(ctx, database.CreatePostParams{Title: postDetails.Title, Link: postDetails.Link, Text: postDetails.Text, IDSubreddit: postDetails.CommunityId, IDUser: postDetails.UserId})
	if err != nil {
		return 0, getServiceErrorConst(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, getServiceErrorConst(err)
	}

	return insertId, ServiceError{Type: NoError}
}
