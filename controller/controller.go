package controller

import (
	"bukumanga-api/config"
	"bukumanga-api/model"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var db *sql.DB

func init() {
	var err error
	conString := config.GetPostgresConnectionString()
	db, err = sql.Open(config.GetDBType(), conString)
	if err != nil {
		panic(err)
	}
}

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	}
}

func GetEntries() echo.HandlerFunc {
    return func(c echo.Context) error {
		rows, err := db.Query("SELECT id, title, url, domain, bookmark_count, image, hotentried_at, published_at FROM entries")
        if err != nil {
			return errors.Wrapf(err, "connot get entries")
        }
        defer rows.Close()

		entries := []model.Entry{}
        for rows.Next() {
			entry := model.Entry{}
            if err := rows.Scan(
				&entry.ID,
				&entry.Title,
				&entry.URL,
				&entry.Domain,
				&entry.BookmarkCount,
				&entry.Image,
				&entry.HotentriedAt,
				&entry.PublishedAt); err != nil {
				log.Fatalln(err)
            }
            entries = append(entries, entry)
        }

        return c.JSON(http.StatusOK, entries)
    }
}
