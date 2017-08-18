package dbobjects

import (
	"time"
)

// ThreadsUsers is auto-generated model of table "threads_users"
type ThreadsUsers struct {

	// ID column: id
	ID *int `db:"id,omitempty"`

	// ThreadID column: thread_id
	ThreadID *int `db:"thread_id,omitempty"`

	// TalkUserID column: talk_user_id
	TalkUserID *int `db:"talk_user_id,omitempty"`

	// Status column: status
	Status *int `db:"status,omitempty"`

	// CreatedAt column: created_at
	CreatedAt *time.Time `db:"created_at,omitempty"`

	// UpdatedAt column: updated_at
	UpdatedAt *time.Time `db:"updated_at,omitempty"`
}
