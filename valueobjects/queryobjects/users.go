package queryobjects

// User レコードとして取得するThread関連情報の１つ
type User struct {
	IstyleID     *string `db:"istyle_id,omitempty"`
	SpecialistID *string `db:"specialist_id,omitempty"`
}
