package quizrunner

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"gophercises/internal/problemposer"
)

type QuizRunner struct {
	problems       map[string]int
	timeoutSeconds time.Duration
}

func NewQuizRunner(problems map[string]int, timeoutSeconds time.Duration) *QuizRunner {
	return &QuizRunner{problems: problems, timeoutSeconds: timeoutSeconds}
}

func (quiz *QuizRunner) Run() map[string]bool {
	problemChannel := make(chan problemposer.Problem)
	answerChannel := make(chan bool)

	problemposer.Initialise(problemChannel, answerChannel)

	waitForUser()

	return poseQuestions(quiz, problemChannel, answerChannel)
}

func poseQuestions(quiz *QuizRunner, problemChannel chan problemposer.Problem, answerChannel chan bool) map[string]bool {
	results := make(map[string]bool)
	for question, answer := range quiz.problems {
		problemChannel <- problemposer.Problem{Question: question, Answer: answer}

		select {
		case response := <-answerChannel:
			results[question] = response
		case <-time.After(quiz.timeoutSeconds * time.Second):
			return results
		}
	}

	return results
}

func waitForUser() {
	fmt.Println("Press enter to start quiz")
	reader := bufio.NewReader(os.Stdin)
	key := make([]byte, 1)
	reader.Read(key)
}
