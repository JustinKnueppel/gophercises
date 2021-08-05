package quiz_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/JustinKnueppel/gophercises/quiz"

	"github.com/google/go-cmp/cmp"
)

func TestReadInput(t *testing.T) {
	tests := map[string]struct {
		reader io.Reader
		want   []quiz.Question
	}{
		"empty": {
			reader: bytes.NewReader([]byte("")),
			want:   []quiz.Question{},
		},
		"one_line": {
			reader: bytes.NewReader([]byte("1+1,2")),
			want:   []quiz.Question{{Question: "1+1", Answer: "2"}},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := quiz.ReadInput(tc.reader)
			if err != nil {
				t.Fatalf("Failed to read input")
			}

			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
