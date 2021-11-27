package notes

import "fmt"

// Select returns the notes that match the given ids.
// If no ids are given, all notes are returned.
func (book *Book) Select(ids ...int) ([]*Note, error) {
	if book == nil {
		return nil, fmt.Errorf("book is nil")
	}

	notes := make([]*Note, 0, len(book.notes))

	if len(ids) == 0 {
		book.log("[SELECT]\tall\n")

		for _, note := range book.notes {
			notes = append(notes, note)
		}
	}

	for _, id := range ids {
		book.log("[SELECT]\tby id:\t%d\n", id)

		note, ok := book.notes[id]
		if !ok {
			err := ErrNotFound(id)
			book.log("[SELECT]\terror:\t%s\n", err)
			return nil, err
		}
		notes = append(notes, note)

	}

	book.log("[SELECT]\tfound:\t%d\n", len(notes))

	return notes, nil
}
