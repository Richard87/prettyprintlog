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

func TimeOnlyOption(w *zerolog.ConsoleWriter) { w.TimeFormat = time.TimeOnly }

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	flag.Parse()

	writer := zerolog.NewConsoleWriter(TimeOnlyOption)
	for scanner.Scan() {
		content := scanner.Text()

		_, err := writer.Write([]byte(content))
		if err != nil {
			fmt.Println(content)
		}
	}
}
