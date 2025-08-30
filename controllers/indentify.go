package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadshanoop/identity-reconciliation/services"
)

func IdentifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UserServiceInstance.ReconcileUser(c)
	}
}
