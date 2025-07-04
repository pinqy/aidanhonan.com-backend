package middleware_test

import (
	"net/http/httptest"
	"testing"

	"aidanhonan.com/backend/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNoRouteHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	middleware.NoRouteHandler()(c)

	expectedErrorBody := `{"code":"FORBIDDEN","message":"Access Denied"}`

	assert.Equal(t, 403, w.Code)
	assert.Equal(t, expectedErrorBody, w.Body.String())
}
