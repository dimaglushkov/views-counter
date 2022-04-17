package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dimaglushkov/views_counter/app"
)

type repository struct {
	*sql.DB
	urls map[string]bool
}

func New(db *sql.DB, urls ...string) app.Repository {
	repo := repository{
		DB:   db,
		urls: make(map[string]bool, len(urls)),
	}

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS views
(
    "url"    	TEXT  NOT NULL PRIMARY KEY,
    "count" 	BIGINT NOT NULL DEFAULT 0
);
`)
	if err != nil {
		panic(fmt.Errorf("error while creating \"views\" table: %v", err))
	}
	for _, url := range urls {
		repo.urls[url] = true
		_, err = db.Exec(`INSERT OR IGNORE INTO views(url) VALUES($1)`, url)
		if err != nil {
			panic(fmt.Errorf("error while inserting \"%s\" record: %v", url, err))
		}
	}

	return repo
}

func (r repository) Visit(ctx context.Context, url string) (int64, error) {
	var res int64
	var err error

	if _, ok := r.urls[url]; !ok {
		return 0, app.UnknownUrlError{Url: url}
	}

	err = r.QueryRowContext(ctx, `SELECT "count" FROM views WHERE "url" = $1`, url).Scan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, app.UnknownUrlError{Url: url}
		}
		return 0, fmt.Errorf("repository error while visiting %s: %v", url, err)
	}
	_, err = r.ExecContext(ctx, `UPDATE views SET "count" = "count" + 1 WHERE "url" = $1;`, url)
	if err != nil {
		return 0, fmt.Errorf("repository error while updating \"%s\" views count: %v", url, err)
	}
	return res + 1, nil
}
