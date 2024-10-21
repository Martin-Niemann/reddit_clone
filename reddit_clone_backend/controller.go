package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

type Controller struct {
	handler http.Handler
}

func (c *Controller) setup() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /c/{name}", func(writer http.ResponseWriter, req *http.Request) {
		println(req.PathValue("name"))
	})

	c.handler = handlers.LoggingHandler(os.Stdout, mux)
}

func (c *Controller) run() {
	log.Fatal(http.ListenAndServe(":8090", c.handler))
}
