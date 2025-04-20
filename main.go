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

	router := gin.Default()
	router.POST("/test", testApi)

	router.Run("localhost:8080")
}

func testApi(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "test response")
}
