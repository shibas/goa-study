package queryobjects

// MessagePack メッセージデータを表現する
type MessagePack struct {
	ID        int     `db:"message_id"`
	UUID      string  `db:"message_uuid"`
	CreatedAt string  `db:"message_created_at,omitempty"`
	UpdatedAt string  `db:"message_updated_at,omitempty"`
	Title     *string `db:"message_title,omitempty"`
	Type      *string `db:"message_type,omitempty"`
	Body      *string `db:"message_body,omitempty"`
	User      `db:",inline"`
	Thumbnail `db:",inline"`
}
