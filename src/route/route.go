package route

import (
	"app/src/middleware"
	"app/src/pkg/conf"

	_ "app/docs"
	"app/src/api"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRoute() *gin.Engine {
	engine := gin.New()

	if conf.Server.RunMode == "release" {
		engine.Use(gin.Logger())
	}
	// engine.Use(gin.Recovery())
	gin.SetMode(conf.Server.RunMode)

	engine.Use(middleware.Logger())

	// swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	loadRoutes(engine)

	return engine
}

func loadRoutes(gin *gin.Engine) {
	gin.POST("/topic/topics/:id", api.Topics)
	gin.POST("/topic/save", api.Save)
}
