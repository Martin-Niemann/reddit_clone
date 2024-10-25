package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
)

type Controller struct {
	handler http.Handler
	server  http.Server
	service Service
}

func (c *Controller) getSubreddit(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	url := req.PathValue("url")

	result, err := c.service.getSubreddit(ctx, url)
	if err != nil {
		writer.Write([]byte(err.Error()))
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	err = json.NewEncoder(writer).Encode(*result)
	if err != nil {
		writer.Write([]byte(err.Error()))
	}
}

func (c *Controller) setupRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /c/{name}", func(writer http.ResponseWriter, req *http.Request) {
		println(req.PathValue("name"))
	})

	mux.HandleFunc("GET /r/{url}", c.getSubreddit)

	return mux
}

func (c *Controller) Setup(mysql MysqlConnectionVariables) error {
	c.service = NewService(mysql)

	mux := c.setupRoutes(http.NewServeMux())

	c.handler = handlers.LoggingHandler(os.Stdout, mux)

	c.server = http.Server{Addr: ":8090", Handler: c.handler}

	return nil
}

func (c *Controller) Run() {
	go func() {
		log.Fatal(c.server.ListenAndServe())
	}()
	log.Printf("HTTP server succesfully started on address: %s.", c.server.Addr)
}

func (c *Controller) Shutdown(timeoutCtx context.Context) {
	c.service.closeConnection()
	c.server.Shutdown(timeoutCtx)
}
