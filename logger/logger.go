package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var Log *zerolog.Logger

func Init() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.TimeOnly}
	logger := zerolog.New(consoleWriter).With().Timestamp().Logger()
	Log = &logger
}
