package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/JustinKnueppel/gophercises/quiz"
)

var (
	csvFile string
)

func init() {
	flag.StringVar(&csvFile, "csv", "problems.csv", "a csv file in the format of 'question,answer'")

	flag.Parse()
}

func main() {
	input, err := os.Open(csvFile)

	if err != nil {
		log.Fatalf("Failed to open input file %s: %v", csvFile, err)
	}

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
