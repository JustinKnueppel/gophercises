package quiz_test

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/JustinKnueppel/gophercises/quiz"

	"github.com/google/go-cmp/cmp"
)

func TestReadInput(t *testing.T) {
	tests := map[string]struct {
		reader io.Reader
		want   []quiz.Question
		err    error
	}{
		"empty": {
			reader: bytes.NewReader([]byte("")),
			want:   []quiz.Question{},
			err:    nil,
		},
		"one_line": {
			reader: bytes.NewReader([]byte("1+1,2")),
			want:   []quiz.Question{{Question: "1+1", Answer: "2"}},
			err:    nil,
		},
		"one_line_too_many_fields": {
			reader: bytes.NewReader([]byte("1+1,2,3")),
			want:   []quiz.Question{},
			err:    &quiz.MarshalQuestionError{},
		},
		"one_line_too_few_fields": {
			reader: bytes.NewReader([]byte("1+1")),
			want:   []quiz.Question{},
			err:    &quiz.MarshalQuestionError{},
		},
		"multiple_valid_lines": {
			reader: bytes.NewReader([]byte("1+1,2\n2+2,4")),
			want:   []quiz.Question{{Question: "1+1", Answer: "2"}, {Question: "2+2", Answer: "4"}},
			err:    nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := quiz.ReadInput(tc.reader)

			if err == nil && tc.err != nil {
				t.Fatalf("Expected error %t, got nil", tc.err)
			}

			if err != nil {
				switch err.(type) {
				case *quiz.MarshalQuestionError:
					if _, ok := (tc.err).(*quiz.MarshalQuestionError); !ok {
						t.Fatalf("Got %v, expected %v", err, tc.err)
					}
				default:
					t.Fatalf("Got unknown error: %v", err)
				}
			}

			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestAskQuestion(t *testing.T) {
	tests := map[string]struct {
		question quiz.Question
	}{
		"empty_question": {
			question: quiz.Question{Question: "", Answer: "2"},
		},
		"normal_question": {
			question: quiz.Question{Question: "1+1", Answer: "2"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var writer bytes.Buffer
			quiz.AskQuestion(&writer, tc.question)

			got := writer.String()

			if !strings.Contains(got, tc.question.Question) {
				t.Fatalf("Question not asked: %v", got)
			}
		})
	}
}

func TestGetResponse(t *testing.T) {
	tests := map[string]struct {
		reader bufio.Reader
		want   string
	}{
		"empty": {
			reader: *bufio.NewReader(bytes.NewBufferString("\n")),
			want:   "",
		},
		"normal": {
			reader: *bufio.NewReader(bytes.NewBufferString("2\n")),
			want:   "2",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := quiz.GetResponse(&tc.reader)

			if err != nil {
				t.Fatalf("Got unexpected error: %v", err)
			}

			diff := cmp.Diff(got, tc.want)

			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
