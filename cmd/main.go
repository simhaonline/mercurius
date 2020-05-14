package main

import (
	"context"
	"fmt"
	"github.com/golangee/http"
	"github.com/golangee/sql"
	"github.com/golangee/uuid"

	"github.com/worldiety/mercurius/service/sms"
)
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/worldiety/mercurius"

func main() {

	db := sql.MustOpen(sql.Opts{
		Driver:       "mysql",
		User:         "root",
		DatabaseName: "mercurius_test",
	})

	sql.MustMigrate(db)

	repos := sql.MustMakeSQLRepositories(db)

	ctx := sql.WithContext(context.Background(), db)
	smsRepo := repos[0].(sms.Repository)
	if err := smsRepo.Create(ctx, uuid.New(), "1234", "hello sms"); err != nil {
		panic(err)
	}

	users, err := smsRepo.FindAll(ctx, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
	fmt.Println("done")

	usr, err := smsRepo.FindById(ctx, uuid.MustParse("353f07bd-9f59-4d67-888d-cc362fed7221"))
	if err != nil {
		panic(err)
	}
	fmt.Println(usr)

	srv := http.NewServer()
	srv.Use(sql.WithTransaction(db))

	ctr, err := http.NewController(srv, sms.NewRestController(smsRepo))
	if err != nil {
		panic(err)
	}
	fmt.Println(ctr)

	err = srv.Start(8080)
	if err != nil {
		panic(err)
	}

}
