package controller

import (
	"bukumanga-api/config"
	"bukumanga-api/model"
	"bukumanga-api/util"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

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
		queryParams := c.QueryParams()
		order := c.QueryParam("order")
		page, _ := strconv.Atoi(c.QueryParam("page"))
		perPage, _ := strconv.Atoi(c.QueryParam("perPage"))

		query := `SELECT * FROM entries
					WHERE
						published_at BETWEEN :startDate AND :endDate AND
						bookmark_count BETWEEN :bookmarkCount AND :bookmarkCountMax`

		input := map[string]interface{}{
			"startDate": queryParams["startDate"],
			"endDate": queryParams["endDate"],
			"bookmarkCount": queryParams["bookmarkCount"],
			"bookmarkCountMax": queryParams["bookmarkCountMax"],
		}

		// ドメイン指定があればフィルターできるようにする
		if len(queryParams["publisherIds"]) > 0 {
			query += ` AND publisher_id IN (:publisherIds)`
			input["publisherIds"] = queryParams["publisherIds"]
		}

		// 値をバインド
		query, args, err := sqlx.Named(query, input)
		if err != nil {
			panic(err)
		}

		// WHERE IN対応
		query, args, err = sqlx.In(query, args...)
		if err != nil {
			panic(err)
		}
		query = db.Rebind(query)

		// キーワードの分割
		keywords := util.TrimSplit(c.QueryParam("keyword"))
		if len(keywords) > 0 {
			query += fmt.Sprintf(" AND (%s)", util.MakeWhereKeywordClause(keywords))
		}

		// 総カウント数を取得
		var count int
		err = db.Get(&count, strings.Replace(query, "*", "COUNT(*)", 1), args...)
		if err != nil {
			panic(err)
		}

		// クエリ実行結果を構造体に格納
		query += util.MakeOrderByClause(order)
		query += util.MakeLimitOffsetClause(page, perPage)
		entries := []model.Entry{}
		err = db.Select(&entries, query, args...)
		if err != nil {
			panic(err)
		}

		// エントリIDリストを取得
		var entryIds []int32;
		for _, entry := range entries {
			entryIds = append(entryIds, entry.ID)
		}

		// 各エントリにコメントとドメイン情報を追加
		publisherMap := getPublisherMap()
		commentMap := getCommentMap(entryIds)
		for i, entry := range entries {
			entries[i].Comments = commentMap[entry.ID]
			entries[i].Publisher = publisherMap[entry.PublisherID]
		}

		// レスポンスを作成
		response := model.Response{Count: count, Entries: entries}

        return c.JSON(http.StatusOK, response)
    }
}

// getCommentMap EntryごとのCommentを取得する
func getCommentMap(entryIds []int32) map[int32][]model.Comment {
	if len(entryIds) == 0 {
		return nil
	}

	query, args, err := sqlx.In(`SELECT * FROM comments WHERE entry_id IN (?) AND rank <= 10 ORDER BY rank`, entryIds)
	if err != nil {
		panic(err)
	}
	query = db.Rebind(query)

	comments := []model.Comment{}
	err = db.Select(&comments, query, args...)
	if err != nil {
		panic(err)
	}

	var commentMap = make(map[int32][]model.Comment, 10)
	for _, comment := range comments {
		commentMap[comment.EntryID] = append(commentMap[comment.EntryID], comment)
	}

	return commentMap
}

// getPublisherMap publisherのマップを取得する
func getPublisherMap() map[int32]model.Publisher {
	publishers := []model.Publisher{}
	err := db.Select(&publishers, `SELECT id, domain, name FROM publishers ORDER BY id`)
	if err != nil {
		panic(err)
	}

	var publisherMap = make(map[int32]model.Publisher)
	for _, publisher := range publishers {
		publisherMap[publisher.ID] = publisher
	}

	return publisherMap
}

// GetPublishers 配信サイト一覧を取得する
func GetPublishers() echo.HandlerFunc {
	return func(c echo.Context) error {
		publishers := []model.Publisher{}
		db.Select(&publishers, `SELECT id, domain, name FROM publishers ORDER BY id`)
		response := model.PublishersResponse{Publishers: publishers}
        return c.JSON(http.StatusOK, response)
    }
}
