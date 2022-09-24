package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "ipData/docs"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	return Router
}
