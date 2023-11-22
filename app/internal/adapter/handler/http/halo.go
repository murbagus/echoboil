package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/murbagus/hexapb-go/pkg/log"
	"github.com/murbagus/hexapb-go/pkg/utils"
	"github.com/murbagus/hexapb-go/pkg/validator"
)

type halo struct {
}

func NewHalo(r *echo.Group) {
	h := &halo{}

	rg := r.Group("/halo")
	rg.GET("", h.Get)
	rg.POST("", h.Post)
}

func (h *halo) Get(c echo.Context) error {
	var err error

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

func (h *halo) Post(c echo.Context) error {
	var err error

	// Untuk mencoba validator
	var st struct {
		Username *string `json:"username" validate:"required" govaf:"username"`
		Password *string `json:"password" validate:"" govaf:"password"`
	}

	if err = c.Bind(&st); err != nil {
		log.ConsoleError(err)

		stat, msg := utils.HttpResponsErrorBinding(err)

		return c.JSON(stat, msg)
	}

	gv := validator.New()
	validator.ValidateStruct(gv, st)
	if gv.HasErrors() {
		return c.JSON(http.StatusBadRequest, gv.ErrJSON())
	}

	return c.JSON(http.StatusOK, st)
}
