package server

import (
	"fmt"
	"net/http"
	config "todo_app/internal/configs"

	"github.com/labstack/echo/v4"
)

type server struct {
	echoImplement *echo.Echo
	port          string
}

type Server interface {
	Run() error
}

func NewServer(config *config.Config) Server {
	e := echo.New()

	return &server{
		echoImplement: e,
		port:          config.Server.Port,
	}
}

func (server *server) Run() error {
	server.echoImplement.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	}) // funcの部分をファイルを返すように変更する

	err := server.echoImplement.Start(server.port)

	if err != nil {
		return fmt.Errorf("calling Start: %w", err)
	}
	return nil
}
