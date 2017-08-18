package dbobjects

import (
	"time"
)

// ThreadMeta is auto-generated model of table "thread_meta"
type ThreadMeta struct {

	// ThreadUserID column: thread_user_id
	ThreadUserID int `db:"thread_user_id"`

	// Title column: title
	Title string `db:"title"`

	// CreatedAt column: created_at
	CreatedAt *time.Time `db:"created_at,omitempty"`

	// UpdatedAt column: updated_at
	UpdatedAt *time.Time `db:"updated_at,omitempty"`
}
