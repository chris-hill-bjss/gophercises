package problemreader

import (
	"log"
	"strconv"
	"strings"
)

func Read(content []byte) map[string]int {
	problems := make(map[string]int)

	for _, v := range strings.Split(string(content), "\n") {
		question, answer := parse(v)
		problems[question] = answer
	}

	return problems
}

func parse(problem string) (string, int) {
	problemParts := strings.Split(strings.TrimSuffix(problem, "\r"), ",")

	question := problemParts[0]
	answer, err := strconv.Atoi(problemParts[1])

	if err != nil {
		log.Panic(err)
	}

	return question, answer
}
