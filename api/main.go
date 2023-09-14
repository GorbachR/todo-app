package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/GorbachR/todo-app/api/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

var (
	dbUrl string = "root:dev@/todo-app?multiStatements=true"
)

func runStartupMigration(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./data/migration", "mysql", driver)

	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func InitDb() *sql.DB {
	var err error
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	runStartupMigration(db)
	return db
}

func main() {

	listendAddr := flag.String("listenAddr", ":3000", "the server address")

	db := InitDb()
	defer db.Close()

	r := gin.Default()
	routes.SetupRoutes(r, db)

	fmt.Println("Server running on port:", *listendAddr)
	log.Fatal(r.Run("127.0.0.1" + *listendAddr))
}
