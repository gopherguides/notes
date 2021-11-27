package notes

import (
	"encoding/json"
	"fmt"
	"io"
)

// A Book allows for management, such as
// inserting, updating, and deleting notes.
type Book struct {
	Logger io.Writer // Logger for verbose logging

	notes map[int]*Note
	curID int
}

// Len returns the number of notes in the book.
func (book *Book) Len() int {
	if book == nil {
		return 0
	}

	return len(book.notes)
}

// Restore restores the book from the provided reader.
// The reader should contain a JSON-encoded book.
func (book *Book) Restore(r io.Reader) error {
	if book == nil {
		return nil
	}

	if r == nil {
		return fmt.Errorf("no reader provided")
	}

	book.log("[RESTORE]\tstarting:\t(%d)\n", len(book.notes))

	err := json.NewDecoder(r).Decode(book)

	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to decode book: %w", err)
	}

	book.log("[RESTORE]\tfinished:\t(%d)\n", len(book.notes))

	return nil
}

// Backup backs up the book, in JSON, to the provided writer.
func (book *Book) Backup(w io.Writer) error {
	if book == nil {
		return nil
	}

	if w == nil {
		return fmt.Errorf("no writer provided")
	}

	book.log("[BACKUP]\tstarting:\t(%d)\n", len(book.notes))

	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	if err := enc.Encode(book); err != nil {
		return fmt.Errorf("failed to encode book: %w", err)
	}

	book.log("[BACKUP]\tfinished:\t(%d)\n", len(book.notes))
	return nil
}

func (book *Book) log(format string, a ...interface{}) {
	if book == nil || book.Logger == nil {
		return
	}

	fmt.Fprintf(book.Logger, format, a...)
}
