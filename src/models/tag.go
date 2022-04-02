package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Tag struct {
	GUID      string    `db:"guid" json:"guid"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt null.Time `db:"updated_at" json:"updated_at"`
	DeletedAt null.Time `db:"deleted_at" json:"deleted_at"`
}
