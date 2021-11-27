package notes

import (
	"encoding/json"
	"fmt"
)

type bookJSON struct {
	CurID int           `json:"cur_id,omitempty"`
	Notes map[int]*Note `json:"notes,omitempty"`
}

// MarshalJSON returns the JSON encoding of the book.
func (book *Book) MarshalJSON() ([]byte, error) {
	if book == nil {
		return nil, fmt.Errorf("book is nil")
	}

	b := bookJSON{
		CurID: book.curID,
		Notes: book.notes,
	}

	return json.MarshalIndent(b, "", "  ")
}

// UnmarshalJSON parses the JSON-encoded data and
// stores the result in the book.
func (book *Book) UnmarshalJSON(data []byte) error {
	if book == nil {
		return fmt.Errorf("book is nil")
	}

	b := &bookJSON{}
	if err := json.Unmarshal(data, b); err != nil {
		return err
	}

	book.curID = b.CurID
	book.notes = b.Notes

	return nil
}
