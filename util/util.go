package util

import (
	"regexp"
	"strings"
)

// TrimSplit 文字列を各ワードについてTrimも行った上でスペース区切りで分割する
func TrimSplit(str string) []string {
	rep := regexp.MustCompile(`\s+`)
	keyword := strings.TrimSpace(str)
	keyword = rep.ReplaceAllString(keyword, " ")
	var keywords []string
	if keyword != "" {
		keywords = strings.Split(keyword, " ")
	}
	return keywords
}
