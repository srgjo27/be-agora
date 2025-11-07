package http

import "github.com/gin-gonic/gin"

func NewRouter(
	userHandler *UserHandler,
	authMiddleware *AuthMiddleware,
	categoryHandler *CategoryHandler,
	threadHandler *ThreadHandler,
	postHandler *PostHandler,
	voteHandler *VoteHandler,
) *gin.Engine {
	router := gin.Default()

	router.Use(SetupCORS())

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
			protected.DELETE("/threads/:thread_id", threadHandler.Delete)
			protected.PATCH("/threads/:thread_id", threadHandler.Update)

			protected.POST("/threads/:thread_id/posts", postHandler.Create)

			protected.POST("/threads/:thread_id/vote", voteHandler.VoteOnThread)
			protected.POST("/posts/:post_id/vote", voteHandler.VoteOnPost)

			postsGroup := api.Group("/threads/:thread_id/posts")
			{
				postsGroup.GET("", postHandler.GetByThreadID)
			}
		}

		api.GET("/categories", categoryHandler.GetAll)

		api.GET("/threads", threadHandler.GetAll)
		api.GET("/threads/:thread_id", threadHandler.GetByID)
	}

	return router
}
