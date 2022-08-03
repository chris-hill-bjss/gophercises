package problemposer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	Question string
	Answer   int
}

func Initialise(problemChannel <-chan Problem, answerChannel chan<- bool) {
	go func() {
		for problem := range problemChannel {
			fmt.Println(problem.Question)
			reader := bufio.NewReader(os.Stdin)

			input, err := reader.ReadString('\n')
			if err != nil {
				log.Panic("failed reading user input")
			}

			guess, ok := sanitiseInput(input)

			answerChannel <- ok && (guess == problem.Answer)
		}
	}()
}

func sanitiseInput(input string) (int, bool) {
	guess, err := strconv.Atoi(strings.TrimSuffix(strings.TrimSpace(input), "\r\n"))

	if err != nil {
		return 0, false
	}

	return guess, true
}
