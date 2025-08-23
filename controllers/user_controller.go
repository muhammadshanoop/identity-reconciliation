package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadshanoop/identity-reconciliation/services"
)

func IdentifyUser(c *gin.Context) {
	services.ReconcileUser(c)
}
