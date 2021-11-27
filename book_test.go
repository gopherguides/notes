package notes

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func Test_Book_Backup_Restore(t *testing.T) {
	t.Parallel()

	b1 := &Book{}

	if b1.Len() != 0 {
		t.Fatalf("book should be empty")
	}

	const N = 5
	for i := 0; i < N; i++ {
		note := &Note{
			Body: "test!",
		}

		err := b1.Insert(note)
		if err != nil {
			t.Fatal(err)
		}
	}

	if b1.Len() != N {
		t.Fatalf("expected %d, got %d", N, b1.Len())
	}

	bb := &bytes.Buffer{}
	err := b1.Backup(bb)
	if err != nil {
		t.Fatal(err)
	}

	act := bb.Bytes()

	b2 := &Book{
		Logger: os.Stdout,
	}

	err = b2.Restore(bytes.NewReader(act))
	if err != nil {
		t.Fatal(err)
	}

	if b2.Len() != N {
		t.Fatalf("expected %d, got %d", N, b2.Len())
	}

	an := fmt.Sprint(b1.notes)
	en := fmt.Sprint(b2.notes)

	if an != en {
		t.Fatalf("expected %s, got %s", an, en)
	}

	if b1.curID != b2.curID {
		t.Fatalf("expected %d, got %d", b1.curID, b2.curID)
	}
}
