package server

import (
	"github.com/gin-gonic/gin"

	"github.com/pisitj/app-push-to-talk/2-code/backend/handler"
	"github.com/pisitj/app-push-to-talk/2-code/backend/middleware"
)

func Initialize() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	{
		router.Use(middleware.CheckUserLogin())

		router.GET("/chat", handler.GetChat)
		router.POST("/chat", handler.CreateChat)
		router.GET("/chat/:chat_id/message", handler.GetChatHistory)
		router.POST("/chat/:chat_id/message", handler.CreateChatMessage)

		userWhiteListRouter := router.Group("")
		{
			userWhiteListRouter.Use(middleware.CheckUserWhitelist())
			userWhiteListRouter.GET("/push-to-talk/chat/:chat_id/message/:message_id", handler.GetPushToTalk)
			userWhiteListRouter.POST("/push-to-talk/chat/:chat_id/message", handler.CreatePushToTalk)
		}

		adminRouter := router.Group("")
		{
			adminRouter.Use(middleware.CheckUserAdmin())
			adminRouter.GET("/push-to-talk/report", handler.GetReportForPushToTalk)
		}
	}

	router.Run(":8080")
}
