package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	flag "github.com/spf13/pflag"
)

func init() {
	flag.StringVarP(&zerolog.TimestampFieldName, "timestamp", "t", "time", "Timestamp field name")
	flag.StringVarP(&zerolog.MessageFieldName, "message", "m", "message", "Message field name")
	flag.StringVarP(&zerolog.LevelFieldName, "level", "l", "level", "Level field name")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	flag.Parse()

	writer := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly}
	for scanner.Scan() {
		content := scanner.Text()

		_, err := writer.Write([]byte(content))
		if err != nil {
			fmt.Println(content)
		}
	}
}
