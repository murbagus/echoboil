package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/murbagus/hexapb-go/pkg/log"
)

type halo struct {
}

func NewHalo(r *echo.Group) {
	h := &halo{}

	rg := r.Group("/halo")
	rg.GET("", h.Get)
}

func (h *halo) Get(c echo.Context) error {
	a, err := strconv.Atoi("asd")
	if err != nil {
		log.FileError(err)
	}

	// Untuk mencoba membuat panic
	// fmt.Println(a)
	// z := "halo"
	z := 10 / a

	return c.String(http.StatusOK, fmt.Sprint(z))
}
