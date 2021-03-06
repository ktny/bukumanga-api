package model

import (
	"database/sql"
	"time"
)

type Entry struct {
	ID int32 `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	URL string `db:"url" json:"url"`
	Domain string `db:"domain" json:"domain"`
	BookmarkCount int16 `db:"bookmark_count" json:"bookmark_count"`
	Image sql.NullString `db:"image" json:"image"`
	PublisherID int32 `db:"publisher_id" json:"publisher_id"`
	HotentriedAt time.Time `db:"hotentried_at" json:"hotentried_at"`
	PublishedAt time.Time `db:"published_at" json:"published_at"`
	Comments []Comment `db:"comments" json:"comments"`
	Publisher Publisher `db:"publisher" json:"publisher"`
	IsTrend bool `db:"is_trend" json:"is_trend"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Comment struct {
	ID int32 `db:"id" json:"id"`
	EntryID int32 `db:"entry_id" json:"entry_id"`
	Rank int16 `db:"rank" json:"rank"`
	Username string `db:"username" json:"username"`
	Icon string `db:"icon" json:"icon"`
	Content string `db:"content" json:"content"`
	CommentedAt time.Time `db:"commented_at" json:"commented_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Publisher struct {
	ID int32 `db:"id" json:"id"`
	Domain string `db:"domain" json:"domain"`
	Name string `db:"name" json:"name"`
	// Icon string `db:"icon" json:"icon"`
	// CreatedAt time.Time `db:"created_at" json:"created_at"`
	// UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Response struct {
	Count int `db:"count" json:"count"`
	Entries []Entry `db:"entries" json:"entries"`
}

type PublishersResponse struct {
	Publishers []Publisher `db:"publishers" json:"publishers"`
}
