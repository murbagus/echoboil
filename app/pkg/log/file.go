package log

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/murbagus/hexapb-go/pkg/dir"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
)

func fileLogger() (zerolog.Logger, *os.File) {
	curdate := time.Now().Format("2006-01-02")
	logpath := dir.GetLogPath()
	logname := fmt.Sprintf("%s/%s.log", logpath, curdate)

	file, err := os.OpenFile(logname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		ConsoleFatal(eris.New(err.Error()))
	}

	return zerolog.New(file).With().Timestamp().Logger(), file
}

// FileError memberikan output error kedalam file log
func FileError(err error) {
	fl, f := fileLogger()

	defer f.Close()

	err = eris.New(err.Error())

	stack := eris.ToJSON(err, true)
	s, _ := json.Marshal(stack)

	fl.Error().RawJSON("error", s).Send()
}
