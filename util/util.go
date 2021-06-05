package util

import "fmt"

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
