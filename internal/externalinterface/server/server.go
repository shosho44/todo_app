package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

type server struct {
	echoImplement *echo.Echo
}

type Server interface {
	Run() error
}

func NewServer() Server {
	e := echo.New()

	return &server{
		echoImplement: e,
	}
}

func Run() error {
	server.echoImplement.GET('/', func(c echo.Context) {
		return c.String(http.StatusOK, "hello world")
	})

}
