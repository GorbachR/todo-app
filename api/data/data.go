package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

var Db *sql.DB

var (
	dbUrl string = "root:dev@/todo-app?multiStatements=true"
)

func init() {
	var err error
	Db, err = sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(Db, &mysql.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./data/migration", "mysql", driver)

	if err != nil {
		log.Fatal(err)
	}

	m.Up()
}
