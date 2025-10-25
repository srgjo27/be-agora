package http

import "github.com/gin-gonic/gin"

func NewRouter(
	userHandler *UserHandler,
	authMiddleware *AuthMiddleware,
	categoryHandler *CategoryHandler,
	threadHandler *ThreadHandler,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
			auth.POST("/refresh", userHandler.Refresh)
			auth.POST("/logout", userHandler.Logout)
		}

		protected := api.Group("")
		protected.Use(authMiddleware.Authenticate())
		{
			users := protected.Group("/users")
			{
				users.GET("/me", userHandler.GetMyProfile)
			}

			admin := protected.Group("")
			admin.Use(authMiddleware.AdminOnly())
			{
				admin.POST("/categories", categoryHandler.Create)
			}

			protected.POST("/threads", threadHandler.Create)
		}

		api.GET("/categories", categoryHandler.GetAll)

		api.GET("/threads", threadHandler.GetAll)
		api.GET("/threads/:id", threadHandler.GetByID)
	}

	return router
}
