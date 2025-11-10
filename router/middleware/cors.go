package middleware

import (
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORS middleware
func CORS() gin.HandlerFunc {
	// Define allowed origins
	originsString := "https://aidanhonan.com,https://test.aidanhonan.com"
	// originsString += ",http://localhost:4200" // Uncomment for local testing
	var allowedOrigins []string
	if originsString != "" {
		allowedOrigins = strings.Split(originsString, ",")
	}

	// return actual middleware handler func
	return func(c *gin.Context) {
		// Function to check if a given origin is allowed
		isOriginAllowed := func(origin string, allowedOrigins []string) bool {
			return slices.Contains(allowedOrigins, origin)
		}

		// Get the Origin header from the request
		origin := c.Request.Header.Get("Origin")

		// Check if the origin is allowed
		if isOriginAllowed(origin, allowedOrigins) {
			// If the origin is allowed, set CORS headers in the response
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Vary", "Origin")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		}

		// Handle preflight OPTIONS requests by aborting with status 204
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Call the next handler
		c.Next()
	}
}
