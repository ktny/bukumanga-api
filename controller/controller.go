package controller

import (
	"bukumanga-api/config"
	"bukumanga-api/model"
	"bukumanga-api/util"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

const DATE_FMT string = "2006-01-02"
var db *sqlx.DB

func init() {
	var err error
	conString := config.GetPostgresConnectionString()
	db, err = sqlx.Open(config.GetDBType(), conString)
	if err != nil {
		panic(err)
	}
}

// Hello ヘルスチェック
func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	}
}

// GetEntries エントリ一覧を取得する
func GetEntries() echo.HandlerFunc {
    return func(c echo.Context) error {
		// クエリパラメータの取得
		startDate, _ := time.Parse(DATE_FMT, c.QueryParam("startDate"))
		endDate, _ := time.Parse(DATE_FMT, c.QueryParam("endDate"))
		bookmarkCount, _ := strconv.Atoi(c.QueryParam("bookmarkCount"))
		bookmarkCountMax, _ := strconv.Atoi(c.QueryParam("bookmarkCountMax"))
		order := c.QueryParam("order")
		page, _ := strconv.Atoi(c.QueryParam("page"))
		perPage, _ := strconv.Atoi(c.QueryParam("perPage"))

		// キーワードの分割
		rep := regexp.MustCompile(`\s+`)
		keyword := strings.TrimSpace(c.QueryParam("keyword"))
		keyword = rep.ReplaceAllString(keyword, " ")
		var keywords []string
		if keyword != "" {
			keywords = strings.Split(keyword, " ")
		}

		// SQLクエリの構築
		query := `SELECT * FROM entries`
		whereClause := util.MakeWhereClause()

		if len(keywords) > 0 {
			whereClause += fmt.Sprintf(" AND (%s)", util.MakeWhereKeywordClause(keywords))
		}

		// 総カウント数を取得
		var count int
		db.Get(&count, `SELECT COUNT(*) FROM entries` + whereClause, startDate, endDate, bookmarkCount, bookmarkCountMax)

		// クエリ実行結果を構造体に格納
		query += whereClause
		query += util.MakeOrderByClause(order)
		query += util.MakeLimitOffsetClause(page, perPage)
		entries := []model.Entry{}
		db.Select(&entries, query, startDate, endDate, bookmarkCount, bookmarkCountMax)
		for i, entry := range entries {
			comments := []model.Comment{}
			db.Select(&comments, `SELECT * FROM comments WHERE entry_id = $1 ORDER BY rank`, entry.ID)
			entries[i].Comments = comments
		}

		// レスポンスを作成
		response := model.Response{Count: count, Entries: entries}

        return c.JSON(http.StatusOK, response)
    }
}
