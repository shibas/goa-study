package dbobjects

import (
	"time"
)

// Thread is auto-generated model of table "threads"
type Thread struct {

	// ID column: id
	ID *int `db:"id,omitempty"`

	// UUID column: uuid
	UUID string `db:"uuid"`

	// CreatedAt column: created_at
	CreatedAt *time.Time `db:"created_at,omitempty"`

	// UpdatedAt column: updated_at
	UpdatedAt *time.Time `db:"updated_at,omitempty"`

	// Status column: status
	Status *int `db:"status"`
}
