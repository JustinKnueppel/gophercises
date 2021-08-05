package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/JustinKnueppel/gophercises/quiz"
)

func main() {
	input := bytes.NewReader([]byte("1+1,2\n2+2,4"))

	questions, err := quiz.ReadInput(input)
	if err != nil {
		log.Fatalf("Failed to read quiz input: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	writer := io.Writer(os.Stdout)

	numCorrect := 0

	for _, question := range questions {
		quiz.AskQuestion(writer, question)

		response, err := quiz.GetResponse(reader)

		if err != nil {
			log.Fatal("Error reading response")
		}

		if quiz.CorrectAnswer(response, question) {
			numCorrect += 1
		}
	}

	fmt.Printf("Quiz finished! You answered %d out of %d questions correctly.\n", numCorrect, len(questions))
}
