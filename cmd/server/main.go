package main

import (
	"net/http"

	auth "gaservice/api/handlers"
	_ "gaservice/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files

//	@title			Your Project API
//	@version		1.0
//	@description	This is a sample server
//	@host			localhost:8080
//	@BasePath		/

func main() {
	r := gin.Default()

	// 一个示例的 API 路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	jwt := r.Group("/auth/jwt")
	jwt.POST("", auth.IssueJWT)
	jwt.GET("/refresh", auth.RefreshJWT)

	// Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
