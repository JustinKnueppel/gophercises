package quiz

import (
	"bufio"
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
	records, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		return []Question{}, err
	}

	questions := []Question{}

	for _, line := range records {
		if len(line) != 2 {
			return []Question{}, &MarshalQuestionError{Line: line}
		}

		question := Question{
			Question: line[0],
			Answer:   line[1],
		}

		questions = append(questions, question)
	}

	return questions, nil
}

func AskQuestion(writer *bufio.Writer, question Question) {
	writer.WriteString(question.Question + "\n")
}
