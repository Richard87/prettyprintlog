package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	flag.String("timestamp", "time", "Timestamp field name")
	flag.String("message", "message", "Message field name")
	flag.String("level", "level", "Level field name")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}

	zerolog.TimestampFieldName = viper.GetString("timestamp")
	zerolog.MessageFieldName = viper.GetString("message")
	zerolog.LevelFieldName = viper.GetString("level")
	writer := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly}
	for scanner.Scan() {
		content := scanner.Text()

		_, err := writer.Write([]byte(content))
		if err != nil {
			fmt.Println(content)
		}
	}
}
