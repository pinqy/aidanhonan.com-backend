package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NoRoute Handler
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusForbidden, gin.H{"code": "FORBIDDEN", "message": "Access Denied"})
	}
}
