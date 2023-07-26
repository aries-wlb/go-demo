package app

import (
	"os"

	"net/http"

	"github.com/google/wire"
	"github.com/uptrace/bunrouter"
	"patrick.com/abroad/app/api"
	"patrick.com/abroad/app/logger"
)

var AppSet = wire.NewSet(wire.Struct(new(App), "*"), api.InitRoutes)

type App struct {
	router *bunrouter.Router
}

func (this *App) Run() {
	logger.Info("Starting server...")
	port := os.Getenv("server.port")
	if port == "" {
		port = ":3000"
	}
	logger.Info("listing at 3000")
	http.ListenAndServe(port, this.router)
}
