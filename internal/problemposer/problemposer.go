package problemposer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Initialise(questionChannel <-chan string, answerChannel chan<- int) {
	go func() {
		for question := range questionChannel {
			fmt.Println(question)
			reader := bufio.NewReader(os.Stdin)

			input, _ := reader.ReadString('\n')
			guess, _ := strconv.Atoi(strings.TrimSuffix(input, "\r\n"))

			answerChannel <- guess
		}
	}()
}
