package main

import (
	"net/http"
	"strconv"
)

func (c *Controller) setupRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /subreddit/{url}", c.getSubreddit)

	mux.HandleFunc("GET /post/{id}", c.getPost)

	return mux
}

func (c *Controller) getSubreddit(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	url := req.PathValue("url")

	result, serviceError := c.service.getSubreddit(ctx, url)
	if serviceError.Type != NoError {
		sendErrorResponse(writer, serviceError)
	}

	sendValidResponse(result, writer)
}

func (c *Controller) getPost(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	idInt64, err := strconv.ParseInt(req.PathValue("id"), 10, 32)
	if err != nil {
		writer.Write([]byte(err.Error()))
	}

	result, serviceError := c.service.getPost(ctx, int32(idInt64))
	if serviceError.Type != NoError {
		sendErrorResponse(writer, serviceError)
	}

	sendValidResponse(result, writer)
}