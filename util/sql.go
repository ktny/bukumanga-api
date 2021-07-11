package util

import (
	"fmt"
	"strings"
)

// MakeWhereKeywordClause SQLクエリのキーワードのWhere句を作成する
func MakeWhereKeywordClause(keywords []string) string {
	slice := make([]string, len(keywords))
	for i, keyword := range keywords {
		q := fmt.Sprintf("title ILIKE '%%%s%%'", keyword)
		slice[i] = q
	}

	return strings.Join(slice, " AND ")
}

// MakeOrderByClause SQLクエリのOrderBy句を作成する
// -bookmark_count -> ORDER BY boomark_count DESC, id DESC
// +bookmark_count -> ORDER BY boomark_count ASC, id DESC
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
