//go:build wireinject
// +build wireinject

package di

import (
	"todo_app/internal/configs"
	"todo_app/internal/externalinterface/server"

	"github.com/google/wire"
)

func InitializeServer(*config.Config) server.Server {
	panic(
		wire.Build(
			server.NewServer,
		),
	)
}
