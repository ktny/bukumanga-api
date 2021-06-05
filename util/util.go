package util

import "fmt"

func ConvertDBOrder(order string) string {
	// 1文字切り出しで取得されるのはbyteのためstringに変換
	// マルチバイトの場合は[]runeにキャストしてから切り出す必要がある
	symbol := string(order[0])
	key := order[1:]
	sort := "ASC"
	switch symbol {
	case "+":
		sort = "ASC"
	case "-":
		sort = "DESC"
	default:
		key = order
	}
	return fmt.Sprintf("%s %s", key, sort)
}
