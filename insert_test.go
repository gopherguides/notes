package notes

import "testing"

func Test_Book_Insert(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		note *Note
		book *Book
		err  bool
	}{
		{name: "all good", note: &Note{Body: "test"}, book: &Book{}},
		{name: "empty book", note: &Note{Body: "test"}, err: true},
		{name: "empty note and book", err: true},
		{name: "empty note", book: &Book{}, err: true},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.book.Insert(tt.note)

			if tt.err {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			exp := 1
			act := tt.book.Len()

			if act != exp {
				t.Errorf("expected %v, got %v", exp, act)
			}

		})
	}
}
