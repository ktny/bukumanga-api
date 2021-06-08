package model

import "database/sql"

type Entry struct {
	ID int64 `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	URL string `db:"url" json:"url"`
	Domain string `db:"domain" json:"domain"`
	BookmarkCount int64 `db:"bookmark_count" json:"bookmark_count"`
	Image sql.NullString `db:"image" json:"image"`
	HotentriedAt string `db:"hotentried_at" json:"hotentried_at"`
	PublishedAt string `db:"published_at" json:"published_at"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}
