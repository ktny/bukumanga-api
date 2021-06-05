package controller

import (
	"bukumanga-api/config"
	"bukumanga-api/model"
	"bukumanga-api/util"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const dateFmt string = "2006-01-02"
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

		startDate, _ := time.Parse(dateFmt, c.QueryParam("startDate"))
		endDate, _ := time.Parse(dateFmt, c.QueryParam("endDate"))
		keyword := c.QueryParam("keyword")
		bookmarkCount, _ := strconv.Atoi(c.QueryParam("bookmarkCount"))
		order := util.ConvertDBOrder(c.QueryParam("order"))

		query := `SELECT id, title, url, domain, bookmark_count, image, hotentried_at, published_at
			FROM entries
			WHERE
				hotentried_at BETWEEN $1 AND $2 AND
				(title ILIKE '%' || $3 || '%' OR domain ILIKE '%' || $3 || '%') AND
				bookmark_count > $4`
		query += fmt.Sprintf(" ORDER BY %s", order)

		rows, err := db.Query(query, startDate, endDate, keyword, bookmarkCount)
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
			// Date部分のみ切り出し
			entry.HotentriedAt = entry.HotentriedAt[:10]
			entry.PublishedAt = entry.PublishedAt[:10]

			// fmt.Printf("%+v\n", entry)
            entries = append(entries, entry)
        }

        return c.JSON(http.StatusOK, entries)
    }
}
