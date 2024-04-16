package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly}
	for scanner.Scan() {
		content := scanner.Text()

		_, err := writer.Write([]byte(content))
		if err != nil {
			fmt.Println(content)
		}
	}
}
