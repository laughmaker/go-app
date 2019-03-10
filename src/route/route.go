package route

import (
	"github.com/laughmaker/go-pkg/conf"

	_ "app/docs"
	"app/src/api"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRoute() *gin.Engine {
	engine := gin.New()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	gin.SetMode(conf.ServerConf.RunMode)

	// engine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))

	// swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// routes
	v := engine.Group(conf.AppConf.Name)
	{
		v.GET("/topic/topics/:id", api.Topics)
	}

	return engine
}
