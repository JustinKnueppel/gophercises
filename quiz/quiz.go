package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
)

type Question struct {
	Question string
	Answer   string
}

type MarshalQuestionError struct {
	Line []string
}

func (e *MarshalQuestionError) Error() string {
	return fmt.Sprintf("Failed to marshal question from %v", e.Line)
}

func ReadInput(reader io.Reader) ([]Question, error) {
	records, err := csv.NewReader(reader).Read()
	if err != nil {
		return []Question{}, nil
	}

	if len(records) != 2 {
		return []Question{}, &MarshalQuestionError{Line: records}
	}

	question := Question{
		Question: records[0],
		Answer:   records[1],
	}

	return []Question{question}, nil
}
