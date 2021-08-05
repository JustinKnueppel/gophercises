package quiz

import "io"

type Question struct {
	Question string
	Answer   string
}

func ReadInput(reader io.Reader) ([]Question, error) {
	return []Question{}, nil
}
