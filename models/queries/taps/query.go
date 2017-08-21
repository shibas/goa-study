package taps

import "time"

// Taps DB objects
type Taps struct {
	ID        *int      `db:"id,omitempty"`
	BeerID    int       `db:"thread_id"`
	CreatedAt time.Time `db:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}
