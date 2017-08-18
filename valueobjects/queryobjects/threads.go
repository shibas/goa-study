package queryobjects

import "time"

// ThreadPack レコードとして取得するThread関連情報の１つ
type ThreadPack struct {
	Thread         `db:",inline"`
	ThreadsUser    `db:",inline"`
	ThreadTalkUser `db:",inline"`
}

// ThreadPackForList レコードとして取得するリスト表示用のThread関連情報
type ThreadPackForList struct {
	Thread         `db:",inline"`
	MessageForList `db:",inline"`
	ThreadsUser    `db:",inline"`
	ThreadTalkUser `db:",inline"`
	LatestEntry    `db:",inline"`
}

// ThreadPackForNewArrival レコードとして取得する新着メッセージ用のThread関連情報
type ThreadPackForNewArrival struct {
	Thread         `db:",inline"`
	MessageForList `db:",inline"`
	ThreadTalkUser `db:",inline"`
	ThreadsUser    `db:",inline"`
	LatestEntry    `db:",inline"`
}

// LatestEntry Messagesが最後にcallされた日時
type LatestEntry struct {
	ThreadUserID    *int       `db:"threads_users_id, omitempty"`
	LastAPICallDate *time.Time `db:"last_api_call_date, omitempty"`
}

// MessageForList メッセージデータを表現する
type MessageForList struct {
	ID             *int       `db:"message_id, omitempty"`
	CreatedAt      *time.Time `db:"message_created_at,omitempty"`
	UpdatedAt      *time.Time `db:"message_updated_at,omitempty"`
	Body           *string    `db:"message_body,omitempty"`
	ThumbnailCount *int       `db:"thumbnails_count, omitempty"`
}

// FlatUserThreadPack レコードとして取得するThread関連情報の１つ
type FlatUserThreadPack struct {
	Thread          `db:",inline"`
	FlatThreadsUser `db:",inline"`
}

// MessageThreadPack メッセージを含むスレッド情報の格納
type MessageThreadPack struct {
	Message        `db:",inline"`
	Thread         `db:",inline"`
	ThreadsUser    `db:",inline"`
	ThreadTalkUser `db:",inline"`
}

// Message メッセージデータを表現する
type Message struct {
	ID                int     `db:"message_id"`
	UUID              string  `db:"message_uuid"`
	CreatedAt         string  `db:"message_created_at,omitempty"`
	UpdatedAt         string  `db:"message_updated_at,omitempty"`
	Title             string  `db:"message_title"`
	Type              string  `db:"message_type"`
	Body              *string `db:"message_body,omitempty"`
	MessageTalkUserID *int    `db:"message_talk_user_id,omitempty"`
	Thumbnail         `db:",inline"`
}

// Thumbnail サムネイル画像の情報を設定
type Thumbnail struct {
	ThumbnailID    *int    `db:"thumbnail_id,omitempty"`
	ThumbnailURL   *string `db:"thumbnail_url,omitempty"`
	ThumbnailOrder *int    `db:"thumbnail_order,omitempty"`
	CreatedAt      *string `db:"thumbnail_created_at,omitempty"`
	UpdatedAt      *string `db:"thumbnail_updated_at,omitempty"`
}

// Thread スレッドデータを表現するstruct
type Thread struct {
	ID        int    `db:"thread_id,omitempty"`
	UUID      string `db:"uuid"`
	Status    string `db:"status"`
	CreatedAt string `db:"thread_created_at,omitempty"`
	UpdatedAt string `db:"thread_updated_at,omitempty"`
}

// ThreadsUser ユーザーとスレッドのひも付きデータを表現するstruct
type ThreadsUser struct {
	ID         *int       `db:"threads_users_id,omitempty"`
	ThreadID   *int       `db:"threads_users_thread_id,omitempty"`
	TalkUserID *int       `db:"threads_users_user_id,omitempty"`
	Status     *int       `db:"threads_users_status,omitempty"`
	CreatedAt  *string    `db:"threads_users_created_at,omitempty"`
	UpdatedAt  *string    `db:"threads_users_updated_at,omitempty"`
	ThreadMeta ThreadMeta `db:",inline"`
}

// FlatThreadsUser ユーザーとスレッドのひも付きデータを表現するstruct
type FlatThreadsUser struct {
	ID             int             `db:"threads_users_id,omitempty"`
	ThreadID       int             `db:"threads_users_thread_id"`
	TalkUserID     int             `db:"threads_users_user_id"`
	Status         int             `db:"threads_users_status,omitempty"`
	CreatedAt      string          `db:"threads_users_created_at,omitempty"`
	UpdatedAt      string          `db:"threads_users_updated_at,omitempty"`
	ThreadMeta     *ThreadMeta     `db:",inline"`
	ThreadTalkUser *ThreadTalkUser `db:",inline"`
}

// ThreadMeta ユーザーに紐づくスレッドのメタデータを表現するstruct
type ThreadMeta struct {
	ThreadUserID *int    `db:"meta_thread_user_id,omitempty"`
	Title        *string `db:"title,omitempty"`
	CreatedAt    *string `db:"meta_created_at,omitempty"`
	UpdatedAt    *string `db:"meta_updated_at,omitempty"`
}

// ThreadTalkUser スレッドから引いた時のTalkUser
type ThreadTalkUser struct {
	ID         *int                  `db:"talk_user_id,omitempty"`
	Status     *int                  `db:"talk_user_status,omitempty"`
	CreatedAt  *string               `db:"user_created_at,omitempty"`
	UpdatedAt  *string               `db:"user_updated_at,omitempty"`
	AtMember   *ThreadUserAtMember   `db:",inline"`
	Specialist *ThreadUserSpecialist `db:",inline"`
}

// ThreadUserAtMember スレッドから引いたTalkAtMember
type ThreadUserAtMember struct {
	IstyleID   *string `db:"istyle_id"`
	TalkUserID *int    `db:"at_member_talk_user_id"`
	CreatedAt  *string `db:"at_member_created_at"`
	UpdatedAt  *string `db:"at_member_updated_at"`
}

// ThreadUserSpecialist スレッドから引いたSpecialist
type ThreadUserSpecialist struct {
	BeautyID     *string `db:"beauty_id"`
	SpecialistID *string `db:"specialist_id"`
	ProfileImage *string `db:"specialist_profile_image"`
	TalkUserID   *int    `db:"specialist_talk_user_id"`
	CreatedAt    *string `db:"specialist_created_at"`
	UpdatedAt    *string `db:"specialist_updated_at"`
}
