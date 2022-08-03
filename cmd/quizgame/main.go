package main

import (
	"log"
	"os"

	"gophercises/internal/problemreader"
	"gophercises/internal/quizrunner"
)

func main() {
	content := readInputFile("input.txt")

	problems := inputToProblems(content)

	quizRunner := quizrunner.NewQuizRunner(problems, 30)

	results := quizRunner.Run()

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

func calculateScore(results map[string]bool) int {
	var correct int
	for _, v := range results {
		if v {
			correct++
		}
	}
	return correct
}
