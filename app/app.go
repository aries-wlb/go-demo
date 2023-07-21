package app

import (
	"os"

	"dipont.com/demo/app/controller"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func NewEngine(userApi *controller.UserApi) *gin.Engine {
	engine := gin.Default()
	api := engine.Group("/api")
	user := api.Group("/user")
	{
		user.GET("/", userApi.Get)
	}
	return engine
}

var AppSet = wire.NewSet(wire.Struct(new(App), "*"), NewEngine)

type App struct {
	Engine *gin.Engine
}

func (this *App) run() {
	port := os.Getenv("server.port")
	if port == "" {
		port = "3000"
	}
	this.Engine.Run("0.0.0.0:" + port)
}
