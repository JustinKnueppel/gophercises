package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
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

func AskQuestion(writer io.Writer, question Question) {
	writer.Write([]byte(question.Question + "\n"))
}

func GetResponse(reader *bufio.Reader) (string, error) {
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.Replace(text, "\n", "", -1)

	return text, nil
}

func CorrectAnswer(response string, question Question) bool {
	return strings.Compare(response, question.Answer) == 0
}
