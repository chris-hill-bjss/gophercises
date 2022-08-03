package main

import (
	"log"
	"os"

	"gophercises/internal/problemposer"
	"gophercises/internal/problemreader"
)

func main() {
	content := readInputFile("input.txt")

	problems := inputToProblems(content)

	results := runQuiz(problems)

	score := calculateScore(results)

	log.Printf("Score %v/%v", score, len(problems))
}

func readInputFile(filename string) []byte {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic(err)
	}

	return content
}

func inputToProblems(content []byte) map[string]int {
	return problemreader.Read(content)
}

func runQuiz(problems map[string]int) map[string]bool {
	questionChannel := make(chan string)
	answerChannel := make(chan int)
	problemposer.Initialise(questionChannel, answerChannel)

	results := make(map[string]bool)
	for question, answer := range problems {
		questionChannel <- question

		guess := <-answerChannel

		results[question] = guess == answer
	}

	return results
}

func calculateScore(results map[string]bool) int {
	var correct int
	for _, v := range results {
		if v {
			correct++
		}
	}
	return correct
}
