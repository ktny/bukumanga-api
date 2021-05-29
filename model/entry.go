package model

import "database/sql"

type Entry struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	URL string `json:"url"`
	Domain string `json:"domain"`
	BookmarkCount int64 `json:"bookmark_count"`
	Image sql.NullString `json:"image"`
	HotentriedAt string `json:"hotentried_at"`
	PublishedAt string `json:"published_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
