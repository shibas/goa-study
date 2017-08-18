package dbobjects

import "time"

// ThreadUserLatestEntries Table struct
type ThreadUserLatestEntries struct {
	ThreadUserID    int        `db:"thread_user_id"`
	LastAPICallDate *time.Time `db:"last_api_call_date,omitempty"`
}
