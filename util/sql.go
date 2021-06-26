package util

import (
	"fmt"
	"strings"
)

// MakeWhereClause SQLクエリのWhere句を作成する
func MakeWhereClause() string {
	return ` WHERE
	((hotentried_at BETWEEN $1 AND $2) OR (published_at BETWEEN $1 AND $2)) AND
	bookmark_count BETWEEN $3 AND $4`
}

// MakeWhereKeywordClause SQLクエリのキーワードのWhere句を作成する
func MakeWhereKeywordClause(keywords []string) string {
	slice := make([]string, len(keywords))
	for i, keyword := range keywords {
		q := fmt.Sprintf("(title ILIKE '%%%s%%' OR domain ILIKE '%%%s%%')", keyword, keyword)
		slice[i] = q
	}

	return strings.Join(slice, " AND ")
}

// MakeOrderByClause SQLクエリのOrderBy句を作成する
func MakeOrderByClause(order string) string {
	// 1文字切り出しで取得されるのはbyteのためstringに変換
	// マルチバイトの場合は[]runeにキャストしてから切り出す必要がある
	symbol := string(order[0])
	switch symbol {
	case "+":
		return fmt.Sprintf(" ORDER BY %s %s, id DESC", order[1:], "ASC")
	case "-":
		return fmt.Sprintf(" ORDER BY %s %s, id DESC", order[1:], "DESC")
	default:
		return fmt.Sprintf(" ORDER BY %s %s, id DESC", order, "ASC")
	}
}

// MakeLimitOffsetClause SQLクエリのLIMITとOFFSET句を作成する
func MakeLimitOffsetClause(page int, perPage int) string {
	return fmt.Sprintf(" LIMIT %d OFFSET %d", perPage, page * perPage)
}
