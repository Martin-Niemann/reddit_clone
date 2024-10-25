package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// https://millhouse.dev/posts/graceful-shutdowns-in-golang-with-signal-notify-context
	// source for how to do graceful shutdown

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql := MysqlConnectionVariables{os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS")}

	var controller Controller
	controller.Setup(mysql)
	controller.Run()

	<-ctx.Done()

	stop()
	log.Println("Shutting down gracefully. Press Ctrl+C again to force shutdown.")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go controller.Shutdown(timeoutCtx)

	select {
	case <-timeoutCtx.Done():
		if timeoutCtx.Err() == context.DeadlineExceeded {
			log.Fatalln("Timeout exceeded. Forcing shutdown.")
		}

		os.Exit(0)
	}
}
