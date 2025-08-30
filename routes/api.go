package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(identifyHandler gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.POST("api/identify", identifyHandler)
	return router
}
