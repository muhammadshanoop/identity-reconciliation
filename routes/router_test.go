package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/muhammadshanoop/identity-reconciliation/routes"
	"github.com/stretchr/testify/require"
)

func TestSetupRouter_PostIdentify_RouteExists(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fake := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": "true"})
	}
	r := routes.SetupRouter(fake)
	req := httptest.NewRequest(http.MethodPost, "/api/identify", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), `"ok":"true"`)
}
