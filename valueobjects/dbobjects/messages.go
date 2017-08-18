package dbobjects

import "time"

// Messages is auto-generated model of table "messages"
type Messages struct {

	// ID column: id
	ID *int `db:"id,omitempty"`

	// ThreadID column: thread_id
	ThreadID int `db:"thread_id"`

	// UUID column: uuid
	UUID string `db:"uuid"`

	// CreatedAt column: created_at
	CreatedAt time.Time `db:"created_at,omitempty"`

	// UpdatedAt column: updated_at
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}

// MessagesTalkUsers is auto-generated model of table "messages_talk_users"
type MessagesTalkUsers struct {

	// MessageID column: message_id
	MessageID int `db:"message_id"`

	// TalkUserID column: talk_user_id
	TalkUserID int `db:"talk_user_id"`

	// CreatedAt column: created_at
	CreatedAt time.Time `db:"created_at,omitempty"`

	// UpdatedAt column: updated_at
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}

// MessageMeta is auto-generated model of table "message_meta"
type MessageMeta struct {

	// MessageID column: message_id
	MessageID int `db:"message_id"`

	// Type column: type
	Type string `db:"type"`

	// CreatedAt column: created_at
	CreatedAt time.Time `db:"created_at,omitempty"`

	// UpdatedAt column: updated_at
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}

// Body is auto-generated model of table "body"
type Body struct {

	// MessageID column: message_id
	MessageID int `db:"message_id"`

	// Body column: body
	Body string `db:"body"`

	// CreatedAt column: created_at
	CreatedAt time.Time `db:"created_at,omitempty"`

	// UpdatedAt column: updated_at
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}

// MessagesThumbnails is auto-generated model of table "messages_thumbnails"
type MessagesThumbnails struct {

	// MessageID column: message_id
	MessageID int `db:"message_id"`

	// ThumbnailID column: thumbnail_id
	ThumbnailID int `db:"thumbnail_id"`

	// Order column: order
	Order int `db:"order"`

	// CreatedAt column: created_at
	CreatedAt time.Time `db:"created_at,omitempty"`

	// UpdatedAt column: updated_at
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}

// Thumbnails is auto-generated model of table "thumbnails"
type Thumbnails struct {

	// ID column: id
	ID int `db:"id,omitempty"`

	// URL column: url
	URL string `db:"url"`

	// CreatedAt column: created_at
	CreatedAt time.Time `db:"created_at,omitempty"`

	// UpdatedAt column: updated_at
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}
