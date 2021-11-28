package notes

import (
	"encoding/json"
	"fmt"
	"sort"
)

type bookJSON struct {
	CurID int     `json:"cur_id,omitempty"`
	Notes []*Note `json:"notes,omitempty"`
}

// MarshalJSON returns the JSON encoding of the book.
func (book *Book) MarshalJSON() ([]byte, error) {
	if book == nil {
		return nil, fmt.Errorf("book is nil")
	}

	key := make([]int, 0, len(book.notes))

	for k := range book.notes {
		key = append(key, k)
	}

	notes := make([]*Note, 0, len(book.notes))
	for _, k := range key {
		notes = append(notes, book.notes[k])
	}

	sort.Slice(notes, func(i, j int) bool {
		return notes[i].ID() < notes[j].ID()
	})

	b := bookJSON{
		CurID: book.curID,
		Notes: notes,
	}

	return json.Marshal(b)
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
	book.notes = map[int]*Note{}

	for _, n := range b.Notes {
		book.notes[n.ID()] = n
	}

	return nil
}
