package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/worldiety/mercurius/service/user/repository/postgresql"
	"github.com/worldiety/sqlm"
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"

type AppCtx struct {
	request *http.Request
}

type App struct {
	db *sql.DB
}

func main() {

	host := "localhost"
	port := 5432
	user := "tschinke"
	password := ""
	dbname := "mercurius"

	psqlInfo := fmt.Sprintf("host='%s' port='%d' user='%s' password='%s' dbname='%s' sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic("failed to connect to db: " + err.Error())
	}

	sqlm.MustMigrate(db, postgresql.Migrations...)

	var pgsUsers postgresql.Querier
	pgsUsers = postgresql.New(db)
	users, err := pgsUsers.ListUsers(context.Background())
	if err != nil {
		panic(err)
	}
	_ = users
	_ = pgsUsers
	_ = &App{db: db}

	http.ListenAndServe(":8080", nil)
}
