package quiz

import (
	"encoding/csv"
	"io"
)

type Question struct {
	Question string
	Answer   string
}

func ReadInput(reader io.Reader) ([]Question, error) {
	records, err := csv.NewReader(reader).Read()
	if err != nil {
		return []Question{}, nil
	}

	question := Question{
		Question: records[0],
		Answer:   records[1],
	}

	return []Question{question}, nil
}
