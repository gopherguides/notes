package notes

import "fmt"

// ErrNotFound is returned when a note is not found.
type ErrNotFound int

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("note %d: not found", int(e))
}
