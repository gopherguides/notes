package notes

import (
	"testing"
)

func Test_Book_Update(t *testing.T) {
	t.Parallel()

	b := &Book{}

	note := &Note{
		Body: "test",
	}

	err := b.Insert(note)
	if err != nil {
		t.Fatal(err)
	}

	note.Body = "updated"

	err = b.Update(note)
	if err != nil {
		t.Fatal(err)
	}

	if note.Body != "updated" {
		t.Fatalf("expected %s, got %s", "updated", note.Body)
	}

	if note.updatedAt.Unix() > note.createdAt.Unix() {
		t.Fatalf("expected %d < %d", note.createdAt.Unix(), note.updatedAt.Unix())
	}

	notes, err := b.Select(note.id)
	if err != nil {
		t.Fatal(err)
	}

	if len(notes) != 1 {
		t.Fatal("should be one note")
	}

	if notes[0].Body != note.Body {
		t.Fatal("note body should be updated")
	}

}
