package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadshanoop/identity-reconciliation/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("api/identify", controllers.IdentifyUser)
	return router
}
