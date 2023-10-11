package main

import (
	nhttp "net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/murbagus/hexapb-go/internal/adapter/handler/http"
	"github.com/murbagus/hexapb-go/pkg/log"
	"github.com/rotisserie/eris"
)

func main() {
	// Init echo
	e := echo.New()

	// Handling panic
	if os.Getenv("APP_ENV") == "development" {
		e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
			StackSize: 4 << 10, // 4 KB
			LogLevel:  0,
			LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
				err = eris.New(err.Error())

				// Pada environment dev stactrace akan
				// ditampilkan dalam respons berbentuk HTML
				c.HTML(nhttp.StatusOK, "<p>"+eris.ToCustomString(err, eris.StringFormat{
					Options: eris.FormatOptions{
						InvertTrace: true, // flag that inverts the stack trace output (top of call stack shown first)
						WithTrace:   true, // flag that enables stack trace output
					},
					MsgStackSep:  "<br>", // separator between error messages and stack frame data
					PreStackSep:  "<br>", // separator at the beginning of each stack frame
					StackElemSep: ": ",   // separator between elements of each stack frame
					ErrorSep:     "\n",   // separator between each error in the chain
				})+"</p>")
				log.ConsoleError(err)
				log.FileError(err)

				return err
			},
		}))
	} else {
		e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
			StackSize: 20 << 10, // 4 KB
			LogLevel:  0,
			LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
				err = eris.New(err.Error())
				log.ConsoleError(err)
				log.FileError(err)

				return err
			},
		}))
	}

	// Register handler
	g := e.Group("/api")
	g1 := g.Group("/v1")

	http.NewHalo(g1)

	log.ConsoleInfo("Aplikasi dalam environment:", os.Getenv("APP_ENV"))
	e.Start(":80")
}
