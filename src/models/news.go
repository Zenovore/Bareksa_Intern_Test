package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type News struct {
	GUID        string    `db:"guid" json:"guid"`
	Title       string    `db:"title" json:"title"`
	Content     string    `db:"content" json:"content"`
	Tags        []Tag     `json:"tags"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   null.Time `db:"updated_at" json:"updated_at"`
	DeletedAt   null.Time `db:"deleted_at" json:"deleted_at"`
	PublishedAt null.Time `db:"published_at" json:"published_at"`
}
