// controller/identify_user_test.go
package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/muhammadshanoop/identity-reconciliation/controllers"
	"github.com/muhammadshanoop/identity-reconciliation/services"
	"github.com/stretchr/testify/assert"
)

type mockUserService struct{}

func (m *mockUserService) ReconcileUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "mocked response"})
}

func TestIdentifyUser(t *testing.T) {
	originalService := services.UserServiceInstance
	services.UserServiceInstance = &mockUserService{}
	defer func() { services.UserServiceInstance = originalService }()

	router := gin.Default()
	router.POST("/api/identify", controllers.IdentifyUser())

	req, _ := http.NewRequest("POST", "/api/identify", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "mocked response")
}
