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
	request *http.Request
	*ee.Ctx
}

func (c *AppCtx) Request() ee.Request {
	return ee.NewRequest(c.request)
}

func (c *AppCtx) Users() user.Repository {
	return user.NewSQLRepository(c.Ctx.SQL())
}

func (c *AppCtx) UsersController() *user.Controller {
	return user.NewController(c.Users())
}

type App struct {
	db *sql.DB
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

	app := &App{db: db}

	http.HandleFunc("/", app.ShowUsers)
	http.ListenAndServe(":8080", nil)
}

func (a *App) ShowUsers(w http.ResponseWriter, r *http.Request) {
	ctx := &AppCtx{Ctx: ee.NewCtx(context.Background(), a.db), request: r}

	limit := ee.OptInt(ctx.Request().Param("limit"))
	if limit == 0 {
		limit = 100
	}
	offset := ee.OptInt(ctx.Request().Param("offset"))

	ctx.UsersController().ShowUsers(limit, offset)

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])

	users, err := ctx.Users().FindAll(100, 0)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	for _, usr := range users {
		fmt.Fprintf(w, "Hello, %s!\n", usr.ID.String())
	}
}
