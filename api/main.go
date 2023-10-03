package main

import (
	"context"
	"flag"
	"github.com/GorbachR/todo-app/api/config"
	"github.com/GorbachR/todo-app/api/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	dbUrl string = "root:dev@/todo_app?multiStatements=true&parseTime=true&loc=UTC"
)

func main() {

	listenAddr := flag.String("listenAddr", ":3000", "the server address")

	db, err := config.InitDb(dbUrl)

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer func() {
		log.Println("Closing db connection...")
		if err := db.Close(); err != nil {
			log.Fatal("Db connection closed:", err)
		}
	}()

	r := router.SetupRouter(db)

	svr := &http.Server{
		Addr:    "127.0.0.1" + *listenAddr,
		Handler: r,
	}
	go func() {
		log.Println("Server running on port:", *listenAddr)
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server...")

	if err := svr.Shutdown(context.Background()); err != nil {
		log.Fatal("Server shutdown:", err)
	}

	log.Println("Server exiting")
}
