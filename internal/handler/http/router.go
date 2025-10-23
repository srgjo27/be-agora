package http

import "github.com/gin-gonic/gin"

func NewRouter(userHandler *UserHandler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
		}
	}

	return router
}
