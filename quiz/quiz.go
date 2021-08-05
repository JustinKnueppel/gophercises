package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

// Question holds the prompt and answer for a given question
type Question struct {
	Prompt string
	Answer string
}

// MarshQuestionError is used when a given input cannot be marshalled into a Question
type MarshalQuestionError struct {
	Line []string
}

func (e *MarshalQuestionError) Error() string {
	return fmt.Sprintf("Failed to marshal question from %v", e.Line)
}

// ReadInput marshals the input into a slice of Questions
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
			Prompt: line[0],
			Answer: line[1],
		}

		questions = append(questions, question)
	}

	return questions, nil
}

// AskQuestion prompts the user to answer the given question
func AskQuestion(writer io.Writer, question Question) {
	writer.Write([]byte(question.Prompt + "\n"))
}

// GetResponse receives input from the user
func GetResponse(reader *bufio.Reader) (string, error) {
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.Replace(text, "\n", "", -1)

	return text, nil
}

// CorrectAnswer checks if the user's answer matches the answer to the question
func CorrectAnswer(response string, question Question) bool {
	return strings.Compare(response, question.Answer) == 0
}
