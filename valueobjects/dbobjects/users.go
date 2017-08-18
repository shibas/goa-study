package dbobjects

// TalkAtMember Table struct
type TalkAtMember struct {
	IstyleID   int    `db:"istyle_id"`
	TalkUserID int    `db:"talk_user_id"`
	CreatedAt  string `db:"created_at,omitempty"`
	UpdatedAt  string `db:"updated_at,omitempty"`
}

// Specialist Table struct
type Specialist struct {
	BeautyID     string `db:"beauty_id"`
	SpecialistID int    `db:"specialist_id"`
	TalkUserID   int    `db:"talk_user_id"`
	ProfileImage string `db:"profile_image"`
	CreatedAt    string `db:"created_at,omitempty"`
	UpdatedAt    string `db:"updated_at,omitempty"`
}

// TalkUser Table struct
type TalkUser struct {
	ID        int    `db:"id,omitempty"`
	CreatedAt string `db:"created_at,omitempty"`
	UpdatedAt string `db:"updated_at,omitempty"`
}
