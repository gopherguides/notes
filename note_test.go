package notes

import "testing"

func Test_Note_Short(t *testing.T) {
	t.Parallel()

	const tooLong = "abcdefghijklmnopqrstuvwxyz"
	const tooLongExp = "abcdefghij"

	table := []struct {
		name string
		body string
		exp  string
	}{
		{name: "long", body: tooLong, exp: tooLongExp},
		{name: "short", body: "1234567890", exp: "1234567890"},
		{name: "empty", body: "", exp: ""},
		{name: "new lines", body: "123\n456", exp: "123\\n456"},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			note := &Note{
				Body: tt.body,
			}

			act := note.Short(10)

			if act != tt.exp {
				t.Fatalf("expected %s, got %s", tt.exp, act)
			}

		})
	}
}
