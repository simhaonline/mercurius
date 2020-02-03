package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/worldiety/mercurius/ee"
	"github.com/worldiety/mercurius/user"
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"

type AppCtx struct {
	*ee.Ctx
}

type App struct {
	db    *sql.DB
	users user.Repository
}

func main() {

	db, err := sql.Open("mysql", "root:@/mercurius_test")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic("failed to connect to db: " + err.Error())
	}

	app := &App{db: db, users: user.NewSQLRepository()}

	http.HandleFunc("/", app.HelloServer)
	http.ListenAndServe(":8080", nil)
}

func (a *App) HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])

	ctx := &AppCtx{Ctx: ee.NewCtx(context.Background(), a.db)}
	users, err := a.users.FindAll(ctx.Ctx, struct {
		Limit  int
		Offset int
	}{Limit: 100, Offset: 0})
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	for _, usr := range users {
		fmt.Fprintf(w, "Hello, %s!\n", usr.ID.String())
	}

}
