package log

import (
	"strings"

	"github.com/rotisserie/eris"
)

// ConsoleInfo memberikan output info kedalam console
func ConsoleInfo(msg ...string) {
	consoleLogger.Info().Msg(strings.Join(msg[:], " "))
}

// ConsoleError memberikan output error kedalam console
func ConsoleError(err error) {
	err = eris.New(err.Error())
	trace := eris.ToString(err, true)

	consoleLogger.Error().Msg("\n" + trace)
}

// ConsoleFatal memberikan output fatal kedalam console dan menghentikan program
func ConsoleFatal(err error) {
	err = eris.New(err.Error())
	trace := eris.ToString(err, true)

	consoleLogger.Fatal().Msg("\n" + trace)
}
