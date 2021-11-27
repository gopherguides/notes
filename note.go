package notes

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

// A Note is a single entry in a book.
type Note struct {
	Body      string `json:"body,omitempty"` // The body of the note.
	id        int
	createdAt time.Time
	updatedAt time.Time
}

// ID returns the ID of the note.
func (note *Note) ID() int {
	if note == nil {
		return 0
	}

	return note.id
}

// Short returns a truncated version of the note body.
// If the body is less than the specified length,
// the whole body is returned.
// Newlines are escaped. "\\n"
func (note *Note) Short(i int) string {
	if note == nil || len(note.Body) == 0 {
		return ""
	}

	short := note.Body
	short = strings.ReplaceAll(short, "\n", "\\n")
	// spl := strings.Split(note.Body, "\n")
	// short := spl[0]

	if len(short) <= i {
		return short
	}
	return short[:i]
}

// CreatedAt returns the time the note was created.
func (note *Note) CreatedAt() time.Time {
	if note == nil {
		return time.Time{}
	}

	return note.createdAt
}

// UpdatedAt returns the time the note was last updated.
func (note *Note) UpdatedAt() time.Time {
	if note == nil {
		return time.Time{}
	}

	return note.updatedAt
}

// String returns a human readable representation of the note.
func (note *Note) String() string {
	if note == nil {
		return ""
	}

	bb := &bytes.Buffer{}

	if len(note.Body) > 0 {
		fmt.Fprintf(bb, "%s\n", note.Body)
	}

	var data []string
	if note.id > 0 {
		data = append(data, fmt.Sprintf("ID: %d", note.id))
	}

	if !note.createdAt.IsZero() {
		data = append(data, fmt.Sprintf("Created: %s", note.createdAt.Format(time.RFC3339)))
	}
	if !note.updatedAt.IsZero() {
		data = append(data, fmt.Sprintf("Updated: %s", note.updatedAt.Format(time.RFC3339)))
	}

	if len(data) > 0 {
		fmt.Fprintln(bb, "-------------------")
		fmt.Fprintf(bb, "%s\n", strings.Join(data, "\n"))
	}

	return bb.String()
}

// IsValid returns true if the note body is present
func (note *Note) IsValid() bool {
	if note == nil {
		return false
	}

	return len(note.Body) > 0
}
