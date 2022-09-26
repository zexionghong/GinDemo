package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "ipData/docs"
	"net/http"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /health [get]
func CheckHealth(g *gin.Context) {
	g.JSON(http.StatusOK, "ok")
}

func InitRouter() *gin.Engine {
	Router := gin.Default()
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", CheckHealth)
	}
	return Router
}
