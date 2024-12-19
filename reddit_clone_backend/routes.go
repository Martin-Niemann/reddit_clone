package main

import (
	"net/http"
	"reddit_clone_backend/database"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) setupRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /subreddit/{url}", c.getSubreddit)
	mux.HandleFunc("GET /post/{id}", c.getPost)
	mux.HandleFunc("POST /login", c.login)
	mux.HandleFunc("POST /logout", c.logout)
	mux.HandleFunc("POST /signup", c.signUp)
	mux.Handle("DELETE /deleteaccount", c.authenticationAndAuthorizationMiddleware(http.HandlerFunc(c.deleteAccount)))
	mux.Handle("POST /comment", c.authenticationAndAuthorizationMiddleware(http.HandlerFunc(c.createComment)))
	mux.Handle("POST /comment/score", c.authenticationAndAuthorizationMiddleware(http.HandlerFunc(c.voteOnComment)))
	mux.Handle("POST /subreddit", c.authenticationAndAuthorizationMiddleware(http.HandlerFunc(c.createCommunity)))
	mux.Handle("POST /post", c.authenticationAndAuthorizationMiddleware(http.HandlerFunc(c.createPost)))

	return mux
}

func (c *Controller) getSubreddit(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	url := req.PathValue("url")
	queries := req.URL.Query()

	var result *Subreddit
	var serviceError ServiceError

	println("these are the query values: ", queries.Get("sort_by"), queries.Get("id"))

	// https://stackoverflow.com/a/77865848
	// "The first case that matches the switch statement is executed"
	switch {
	case queries.Get("sort_by") == "newest" && queries.Get("id") != "":
		postId, err := strconv.Atoi(queries.Get("id"))
		if err != nil {
			sendErrorResponse(writer, ServiceError{Type: InvalidArgument})
			return
		}
		result, serviceError = c.service.getSubredditSortedByNewestKeysetPaginated(ctx, url, int32(postId))
	case queries.Get("sort_by") == "newest":
		result, serviceError = c.service.getSubredditSortedByNewestStart(ctx, url)
	case queries.Get("sort_by") == "score" && queries.Get("direction") == "falling" && queries.Get("offset") != "":
		offset, err := strconv.Atoi(queries.Get("offset"))
		if err != nil {
			sendErrorResponse(writer, ServiceError{Type: InvalidArgument})
			return
		}
		result, serviceError = c.service.getSubredditSortedByScoreHighestOffsetPaginated(ctx, url, int32(offset))
	default:
		println("you wrote something wrong dummy!")
	}

	if serviceError.Type != NoError {
		sendErrorResponse(writer, serviceError)
		return
	}

	sendJsonResponse(result, 0, writer)
}

func (c *Controller) getPost(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	idInt64, err := strconv.ParseInt(req.PathValue("id"), 10, 32)
	if err != nil {
		sendErrorResponse(writer, ServiceError{Type: InvalidInput})
		return
	}

	result, serviceError := c.service.getPost(ctx, int32(idInt64))
	if serviceError.Type != NoError {
		sendErrorResponse(writer, serviceError)
		return
	}

	sendJsonResponse(result, 0, writer)
}

func (c *Controller) login(writer http.ResponseWriter, req *http.Request) {
	var loginDetails *LoginDetails = &LoginDetails{}

	err := tryParseJson(req, loginDetails)
	if err != nil {
		sendErrorResponse(writer, ServiceError{Type: InvalidInput})
		return
	}

	ctx := req.Context()
	dbUser, getUserErr := c.service.queries.GetUserByEmail(ctx, loginDetails.Email)

	// https://github.com/OWASP/Go-SCP/blob/master/dist/go-webapp-scp.pdf
	// (page 26) if there is no user with the given email, we continue
	// execution and let bcrypt.CompareHashAndPassword fail
	var expectedHashedPassword string
	if getUserErr == nil {
		expectedHashedPassword = dbUser.HashedPassword
	}

	match := bcrypt.CompareHashAndPassword([]byte(expectedHashedPassword), []byte(loginDetails.Password))
	if match != nil {
		sendErrorResponse(writer, ServiceError{Type: InvalidEmailAndOrPassword})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"username": dbUser.UserName,
		"id":       dbUser.IDUser,
		"iat":      time.Now().UnixMilli(),
	})

	tokenString, _ := token.SignedString(c.jwtSigningKey)

	sendJsonResponse(TokenCookie{Token: tokenString, Expires: time.Now().Add(time.Hour * 24).UnixMilli()}, 200, writer)
}

func (c *Controller) signUp(writer http.ResponseWriter, req *http.Request) {
	var signUpDetails *SignUpDetails = &SignUpDetails{}

	err := tryParseJson(req, signUpDetails)
	if err != nil {
		sendErrorResponse(writer, ServiceError{Type: InvalidInput})
		return
	}

	err = c.validator.Struct(signUpDetails)
	if err != nil {
		validationErrors := createValidationErrorsStruct(err.(validator.ValidationErrors))
		sendErrorResponse(writer, ServiceError{Type: InvalidArgument, ValidationErrors: validationErrors})
		return
	}

	ctx := req.Context()
	result, err := c.service.queries.CheckForExistingUser(ctx, database.CheckForExistingUserParams{Email: signUpDetails.Email, UserName: signUpDetails.UserName})

	if result != false {
		sendErrorResponse(writer, ServiceError{Type: UserAlreadyExists})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpDetails.Password), bcrypt.DefaultCost)
	if err != nil {
		sendErrorResponse(writer, ServiceError{Type: UnexpectedError})
		return
	}

	insertResult, err := c.service.queries.CreateUser(ctx, database.CreateUserParams{UserName: signUpDetails.UserName, Email: signUpDetails.Email, HashedPassword: string(hashedPassword)})
	if err != nil {
		sendErrorResponse(writer, ServiceError{Type: UnexpectedError})
		return
	}

	userId, err := insertResult.LastInsertId()

	sendJsonResponse(UserId{Id: int32(userId)}, 201, writer)
}

func (c *Controller) deleteAccount(writer http.ResponseWriter, req *http.Request) {
	// if we are here, then the submitted jwt token has been verified.
	// as the token is tamper-proof, and it contains the user's id,
	// we can safely use said id to expedite the users wishes.
	ctx := req.Context()

	id, ok := ctx.Value("id").(float64)
	if ok == false {
		sendErrorResponse(writer, ServiceError{Type: UnexpectedError})
		return
	}

	_, err := c.service.queries.DeleteUser(ctx, int32(id))
	if err != nil {
		sendErrorResponse(writer, ServiceError{Type: UserDoesntExist})
		return
	}

	c.logout(writer, req)
}

func (c *Controller) createComment(writer http.ResponseWriter, req *http.Request) {
	// if we are here, then the submitted jwt token has been verified.
	// as the token is tamper-proof, and it contains the user's id,
	// we can safely use said id to expedite the users wishes.
	ctx := req.Context()

	id, ok := ctx.Value("id").(float64)
	if ok == false {
		sendErrorResponse(writer, ServiceError{Type: UnexpectedError})
		return
	}

	var commentDetails *CommentDetails = &CommentDetails{UserId: int32(id)}

	err := tryParseJson(req, commentDetails)
	if err != nil {
		sendErrorResponse(writer, ServiceError{Type: InvalidInput})
		return
	}

	err = c.validator.Struct(commentDetails)
	if err != nil {
		validationErrors := createValidationErrorsStruct(err.(validator.ValidationErrors))
		sendErrorResponse(writer, ServiceError{Type: InvalidArgument, ValidationErrors: validationErrors})
		return
	}

	insertId, serviceError := c.service.addComment(ctx, commentDetails)
	if serviceError.Type != NoError {
		sendErrorResponse(writer, serviceError)
		return
	}

	result, serviceError := c.service.getComment(ctx, int32(insertId))
	if serviceError.Type != NoError {
		sendErrorResponse(writer, serviceError)
		return
	}

	sendJsonResponse(result, 201, writer)
}

func (c *Controller) voteOnComment(writer http.ResponseWriter, req *http.Request) {
	// if we are here, then the submitted jwt token has been verified.
	// as the token is tamper-proof, and it contains the user's id,
	// we can safely use said id to expedite the users wishes.
	ctx := req.Context()

	id, ok := ctx.Value("id").(float64)
	if ok == false {
		sendErrorResponse(writer, ServiceError{Type: UnexpectedError})
		return
	}

	var scoreDetails *ScoreDetails = &ScoreDetails{UserId: int32(id)}

	err := tryParseJson(req, scoreDetails)
	if err != nil {
		sendErrorResponse(writer, ServiceError{Type: InvalidInput})
		return
	}

	_, serviceError := c.service.voteOnComment(ctx, scoreDetails)
	if serviceError.Type != NoError {
		sendErrorResponse(writer, serviceError)
		return
	}
	sendJsonResponse("", 204, writer)
}

func (c *Controller) createCommunity(writer http.ResponseWriter, req *http.Request) {
	// if we are here, then the submitted jwt token has been verified.
	// as the token is tamper-proof, and it contains the user's id,
	// we can safely use said id to expedite the users wishes.
	ctx := req.Context()

	id, ok := ctx.Value("id").(float64)
	if ok == false {
		sendErrorResponse(writer, ServiceError{Type: UnexpectedError})
		return
	}

	var subredditDetails *SubredditDetails = &SubredditDetails{UserId: int32(id)}

	err := tryParseJson(req, subredditDetails)
	if err != nil {
		sendErrorResponse(writer, ServiceError{Type: InvalidInput})
		return
	}

	_, serviceError := c.service.addCommunity(ctx, subredditDetails)
	if serviceError.Type != NoError {
		sendErrorResponse(writer, serviceError)
		return
	}
	sendJsonResponse("", 201, writer)
}

func (c *Controller) createPost(writer http.ResponseWriter, req *http.Request) {
	// if we are here, then the submitted jwt token has been verified.
	// as the token is tamper-proof, and it contains the user's id,
	// we can safely use said id to expedite the users wishes.
	ctx := req.Context()

	id, ok := ctx.Value("id").(float64)
	if ok == false {
		sendErrorResponse(writer, ServiceError{Type: UnexpectedError})
		return
	}

	var postDetails *PostDetails = &PostDetails{UserId: int32(id)}

	err := tryParseJson(req, postDetails)
	if err != nil {
		println(err.Error())
		sendErrorResponse(writer, ServiceError{Type: InvalidInput})
		return
	}

	postId, serviceError := c.service.addPost(ctx, postDetails)
	if serviceError.Type != NoError {
		sendErrorResponse(writer, serviceError)
		return
	}
	sendJsonResponse(PostId{Id: postId}, 201, writer)
}

// https://github.com/OWASP/Go-SCP/blob/master/src/session-management/session.go
func (c *Controller) logout(res http.ResponseWriter, req *http.Request) {
	deleteCookie := http.Cookie{
		Name:    "Auth",
		Value:   "none",
		Expires: time.Now(),
	}

	http.SetCookie(res, &deleteCookie)
}
