package main

import (
	"fx-my-identity/internal/db"
	"fx-my-identity/internal/server"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			server.NewMux,
			db.NewDB,
		),
		fx.Invoke(
			server.StartServer,
		),
	)
	app.Run()
}
