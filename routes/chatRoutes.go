package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetupChatCompletionRoutes(r *gin.Engine) {
	chatGroup := r.Group("/chat")
	{
		chatGroup.POST("/", controllers.ChatCompletion)
		chatGroup.POST("/stream", controllers.StreamChatCompletion)
		chatGroup.POST("/summarize", controllers.SummarizeChatCompletion)
	}
}
