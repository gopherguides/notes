package notes

import "testing"

func Test_Book_Delete(t *testing.T) {
	t.Parallel()

	b := &Book{}

	note := &Note{
		Body: "hello",
	}

	err := b.Insert(note)
	if err != nil {
		t.Fatal(err)
	}

	notes, err := b.Select(note.ID())
	if err != nil {
		t.Fatal(err)
	}

	if len(notes) != 1 {
		t.Fatalf("expected 1 note, got %d", len(notes))
	}

	n2 := notes[0]
	if n2.ID() != note.ID() {
		t.Fatalf("expected note id %d, got %d", note.ID(), n2.ID())
	}

	err = b.Delete(note.ID())
	if err != nil {
		t.Fatal(err)
	}

}
