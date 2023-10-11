package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	consoleLogger zerolog.Logger
)

func init() {
	consoleLoggerOut := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	consoleLogger = zerolog.New(consoleLoggerOut).With().Timestamp().Logger()
}
