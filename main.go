package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strings"
	"time"

	"aidanhonan.com/backend/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	gin_mode := utils.GetEnv("GIN_MODE", gin.ReleaseMode)
	gin.SetMode(gin_mode)
	fmt.Printf("Starting gin server in %s mode\n", strings.ToUpper(gin_mode))

	// Setup router
	router := gin.Default()

	// Use CORS middleware to handle browser pre-flight requests
	router.Use(corsMiddleware())

	// Route handlers
	router.POST("/test", testApi)

	// Start server
	port := utils.GetEnv("PORT", "10000")
	fmt.Printf("Server listening on http://localhost:%s/\n", port)
	if err := router.Run(":" + port); err != nil {
		fmt.Print(fmt.Errorf("failed to start server: %w", err))
	}
}

// API Routes
func testApi(c *gin.Context) {
	randWait := rand.IntN(4) + 1
	time.Sleep((time.Duration(randWait) * time.Second))
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("test response (after %ds)", randWait))
}

// CORS Middleware
func corsMiddleware() gin.HandlerFunc {
	// Define allowed origins
	originsString := "https://aidanhonan.com,http://localhost:4200"
	var allowedOrigins []string
	if originsString != "" {
		allowedOrigins = strings.Split(originsString, ",")
	}

	// return actual middleware handler func
	return func(c *gin.Context) {
		// Function to check if a given origin is allowed
		isOriginAllowed := func(origin string, allowedOrigins []string) bool {
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					return true
				}
			}
			return false
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
