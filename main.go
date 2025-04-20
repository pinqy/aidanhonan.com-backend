package main

import (
	"fmt"
	"net/http"
	"strings"

	"aidanhonan.com/backend/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	gin_mode := utils.GetEnv("GIN_MODE", gin.ReleaseMode)
	gin.SetMode(gin_mode)
	fmt.Printf("Starting gin server in %s mode\n", strings.ToUpper(gin_mode))

	// Setup router
	router := gin.Default()
	router.POST("/test", testApi)

	// Start server
	port := utils.GetEnv("PORT", "10000")
	fmt.Printf("Server listening on http://localhost:%s/\n", port)
	router.Run(":" + port)
}

func testApi(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "test response")
}
