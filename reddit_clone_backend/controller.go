package main

import (
	"context"
	"encoding/json"
	"fmt"
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

func sendValidResponse[T any](result T, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(result)
	if err != nil {
		log.Println(err.Error())

		http.Error(writer, err.Error(), 500)
	}
}

func sendErrorResponse(writer http.ResponseWriter, serviceError ServiceError) {
	switch serviceError.Type {
	case NoResult:
		http.Error(writer, http.StatusText(400), 400)
		return
	case UnexpectedError:
		http.Error(writer, http.StatusText(500), 500)
		return
	default:
		panic(fmt.Sprintf("unexpected main.ServiceError: %+v", serviceError))
	}
}
