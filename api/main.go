package main

import (
	"flag"
	"fmt"
	"github.com/GorbachR/todo-app/api/config"
	"github.com/GorbachR/todo-app/api/router"
	"log"
)

var (
	dbUrl string = "root:dev@/todo_app?multiStatements=true"
)

func main() {

	listendAddr := flag.String("listenAddr", ":3000", "the server address")

	db, err := config.InitDb(dbUrl)

	if err != nil {
		log.Fatal(1)
	}
	defer db.Close()

	migrateStr, err := config.CreateMigrations(db, "file://./data/migration")
	if err != nil {
		log.Fatal(1)
	}

	err = migrateStr.MigrateUp(true)
	if err != nil {
		log.Fatal(1)
	}

	r := router.SetupRouter(db)

	fmt.Println("Server running on port:", *listendAddr)
	log.Fatal(r.Run("127.0.0.1" + *listendAddr))
}
