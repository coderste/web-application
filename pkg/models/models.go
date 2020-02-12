package models

import (
	"errors"
	"time"
)

// ErrNoRecord is the error returned when no records are found in the database
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet is the structure of a snippet inside the database
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
