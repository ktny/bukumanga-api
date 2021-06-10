package util

import "fmt"

// MakeWhereClause SQLクエリのWhere句を作成する
func MakeWhereClause() string {
	return ` WHERE
		hotentried_at BETWEEN $1 AND $2 AND
		(title ILIKE '%' || $3 || '%' OR domain ILIKE '%' || $3 || '%') AND
		bookmark_count >= $4`
}

// MakeOrderByClause SQLクエリのOrderBy句を作成する
func MakeOrderByClause(order string) string {
	// 1文字切り出しで取得されるのはbyteのためstringに変換
	// マルチバイトの場合は[]runeにキャストしてから切り出す必要がある
	symbol := string(order[0])
	switch symbol {
	case "+":
		return fmt.Sprintf(" ORDER BY %s %s", order[1:], "ASC")
	case "-":
		return fmt.Sprintf(" ORDER BY %s %s", order[1:], "DESC")
	default:
		return fmt.Sprintf(" ORDER BY %s %s", order, "ASC")
	}
}

// MakeLimitOffsetClause SQLクエリのLIMITとOFFSET句を作成する
func MakeLimitOffsetClause(page int, perPage int) string {
	return fmt.Sprintf(" LIMIT %d OFFSET %d", perPage, page * perPage)
}
