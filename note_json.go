package notes

import (
	"encoding/json"
	"time"
)

type noteJSON struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MarshalJSON returns the JSON encoding of the note.
func (note *Note) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(noteJSON{
		ID:        note.id,
		Body:      note.Body,
		CreatedAt: note.createdAt,
		UpdatedAt: note.updatedAt,
	}, "", "  ")
}

// UnmarshalJSON parses the JSON-encoded data and
// stores the result in the note.
func (note *Note) UnmarshalJSON(data []byte) error {
	var n noteJSON
	err := json.Unmarshal(data, &n)
	if err != nil {
		return err
	}

	note.id = n.ID
	note.Body = n.Body
	note.createdAt = n.CreatedAt
	note.updatedAt = n.UpdatedAt

	return nil
}
