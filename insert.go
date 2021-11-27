package notes

import (
	"fmt"
	"strings"
	"time"
)

// Insert adds the given notes to the book.
func (book *Book) Insert(notes ...*Note) error {
	if book == nil {
		return fmt.Errorf("book is nil")
	}

	for _, note := range notes {
		if !note.IsValid() {
			return fmt.Errorf("note is invalid")
		}

		book.curID = book.curID + 1

		note.id = book.curID

		now := time.Now()
		note.createdAt = now
		note.updatedAt = now
		note.Body = strings.TrimSpace(note.Body)

		book.log("[INSERT]\tinserting:\t%d\n", note.id)

		if book.notes == nil {
			book.notes = map[int]*Note{}
		}

		book.notes[note.id] = note

		book.log("[INSERT]\tinserted:\t%d\n", note.id)
	}
	return nil
}
