package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/JustinKnueppel/gophercises/quiz"
)

func main() {
	input := bytes.NewReader([]byte("1+1,2\n2+2,4"))

	quiz, err := quiz.ReadInput(input)
	if err != nil {
		log.Fatalf("Failed to read quiz input: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)

	numCorrect := 0

	for _, question := range quiz {
		fmt.Println(question.Question)

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read text input: %v", err)
		}
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare(text, question.Answer) == 0 {
			numCorrect += 1
		}
	}

	fmt.Printf("Quiz finished! You answered %d out of %d questions correctly.\n", numCorrect, len(quiz))
}
