package router

import (
	"NSK_mayoralty_app/controllers"
	"NSK_mayoralty_app/middlewares"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/auth")
	{
		auth.POST("/registration", controllers.Registration)
		auth.POST("/login", controllers.Login)
	}

	points := r.Group("/points")
	{
		points.GET("", controllers.GetPoints)
		points.GET("/:sessionId", controllers.GetSessionsPoints)
		points.POST("", controllers.AddPointToSession)
		points.DELETE("/:pointId", controllers.RemovePointFromSession)
	}
	sessions := r.Group("/sessions")
	sessions.Use(middlewares.AuthMiddleware())
	{
		sessions.GET("/start", controllers.StartSession)
		sessions.GET("/stop", controllers.StopSession)
	}
	swagger := r.Group("/swagger")
	swagger.Use(middlewares.CorsUse)
	{
		swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return r
}
