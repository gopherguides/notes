package notes

import (
	"fmt"
	"time"
)

// Update updates the given notes in the book.
func (book *Book) Update(notes ...*Note) error {
	if book == nil {
		return fmt.Errorf("book is nil")
	}

	for _, note := range notes {
		if !note.IsValid() {
			return fmt.Errorf("note is invalid")
		}

		if note.id <= 0 {
			return fmt.Errorf("note id is invalid %d", note.id)
		}

		if note.createdAt.IsZero() {
			return fmt.Errorf("note created at is invalid")
		}

		book.log("[UPDATE]\tupdating:\t%d\n", note.id)

		if _, ok := book.notes[note.id]; !ok {
			err := ErrNotFound(note.id)
			book.log("[UPDATE]\terror:\t%s\n", err)
			return err
		}

		note.updatedAt = time.Now()
		book.notes[note.id] = note

		book.log("[UPDATE]\tupdated:\t%d\n", note.id)
	}

	return nil
}
