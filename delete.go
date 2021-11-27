package notes

import "fmt"

// Delete removes the given notes from the book, if they exist.
func (book *Book) Delete(ids ...int) error {
	if book == nil {
		return fmt.Errorf("book is nil")
	}

	if len(ids) == 0 {
		return nil
	}

	for _, id := range ids {
		book.log("[DELETE]\tby id:\t%d\n", id)

		_, ok := book.notes[id]
		if !ok {
			err := ErrNotFound(id)
			book.log("[DELETE]\terror:\t%s\n", err)
			return err
		}
		delete(book.notes, id)

	}

	return nil
}
