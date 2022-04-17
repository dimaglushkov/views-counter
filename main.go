package main

import (
	"database/sql"
	"fmt"
	"github.com/dimaglushkov/views_counter/app"
	"github.com/dimaglushkov/views_counter/repositories/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
)

var urls = []string{
	"github.com/dimaglushkov/dimaglushkov",
}

func main() {
	db, err := sql.Open("sqlite3", os.Getenv("DB_DSN"))
	if err != nil {
		panic(fmt.Sprintf("error while opening db connection: %v", err))
	}
	defer db.Close()

	err = http.ListenAndServe(":"+os.Getenv("PORT"), app.NewHandler(sqlite.New(db, urls...)))
	if err != nil {
		panic(err)
	}

}
