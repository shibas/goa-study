package dbobjects

// MuteUser Table struct
type MuteUser struct {
	SourceUserID int     `db:"source_user_id"`
	TargetUserID int     `db:"target_user_id"`
	CreatedAt    *string `db:"created_at,omitempty"`
	UpdatedAt    *string `db:"updated_at,omitempty"`
}
