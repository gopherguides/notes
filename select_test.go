package notes

import (
	"fmt"
	"testing"
)

func Test_Book_Select(t *testing.T) {
	t.Parallel()

	book := &Book{}

	const N = 10

	for i := 0; i < N; i++ {
		err := book.Insert(&Note{
			Body: fmt.Sprintf("Note %d", i+1),
		})

		if err != nil {
			t.Fatal(err)
		}
	}

	if book.Len() != N {
		t.Fatalf("expected %d, got %d", 5, book.Len())
	}

	table := []struct {
		name string
		book *Book
		ids  []int
		err  bool
	}{
		{name: "all good", book: book, ids: []int{1, 2, 3, 4, 5}},
		{name: "empty book", ids: []int{1}, err: true},
		{name: "missing id", book: book, ids: []int{1, 2, 42}, err: true},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			notes, err := tt.book.Select(tt.ids...)
			if tt.err {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}

			if len(notes) != len(tt.ids) {
				t.Fatalf("expected %d, got %d", len(tt.ids), len(notes))
			}

			for i, id := range tt.ids {
				if notes[i].ID() != id {
					t.Fatalf("expected %d, got %d", id, notes[i].ID())
				}
			}

		})
	}
}
