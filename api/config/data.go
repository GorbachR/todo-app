package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

var (
	dbUrl   string = "root:dev@/todo-app?multiStatements=true"
	QueryDb *Queries
)

type Queries struct {
	Db *sql.DB
}

func newQueries(db *sql.DB) *Queries {
	return &Queries{
		Db: db,
	}
}

func InitDbConnection() {
	var err error
	Db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	QueryDb = newQueries(Db)

	driver, err := mysql.WithInstance(Db, &mysql.Config{})

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
