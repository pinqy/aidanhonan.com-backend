package router

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strings"
	"time"

	"aidanhonan.com/backend/router/middleware"
	"aidanhonan.com/backend/utils"
	"github.com/gin-gonic/gin"
)

func Start() {
	gin_mode := utils.GetEnv("GIN_MODE", gin.ReleaseMode)
	gin.SetMode(gin_mode)
	fmt.Printf("Starting gin server in %s mode\n", strings.ToUpper(gin_mode))

	// Setup router
	router := gin.Default()

	// Use CORS middleware to handle browser pre-flight requests
	router.Use(middleware.CORS())

	// Fallback handler for invalid requests
	router.NoRoute(middleware.NoRouteHandler())

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
