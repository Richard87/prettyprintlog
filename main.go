package main

import (
	"bufio"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly}
	for scanner.Scan() {
		content := scanner.Text()

		_, _ = writer.Write([]byte(content))
	}
}
