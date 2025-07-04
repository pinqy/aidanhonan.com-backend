package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"aidanhonan.com/backend/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCORS_PreflightRequestAllowed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(middleware.CORS())

	req, err := http.NewRequest("OPTIONS", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Access-Control-Request-Headers", "content-type")
	req.Header.Add("Access-Control-Request-Method", "POST")
	req.Header.Add("Origin", "https://aidanhonan.com")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert expected CORS preflight response and headers
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Equal(t, "https://aidanhonan.com", rr.Result().Header.Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "Origin", rr.Result().Header.Get("Vary"))
	assert.Equal(t, "true", rr.Result().Header.Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With", rr.Result().Header.Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "POST", rr.Result().Header.Get("Access-Control-Allow-Methods"))
}

func TestCORS_PreflightRequestNotAllowed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(middleware.CORS())

	req, err := http.NewRequest("OPTIONS", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Access-Control-Request-Headers", "content-type")
	req.Header.Add("Access-Control-Request-Method", "POST")
	req.Header.Add("Origin", "https://blocked.com")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert expected CORS preflight response and headers
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Equal(t, "", rr.Result().Header.Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "", rr.Result().Header.Get("Vary"))
	assert.Equal(t, "", rr.Result().Header.Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, "", rr.Result().Header.Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "", rr.Result().Header.Get("Access-Control-Allow-Methods"))
}

func TestCORS_ApiCall(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(middleware.CORS())
	router.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "test response")
	})

	req, err := http.NewRequest("POST", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Access-Control-Request-Headers", "content-type")
	req.Header.Add("Access-Control-Request-Method", "POST")
	req.Header.Add("Origin", "https://aidanhonan.com")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert expected CORS headers
	assert.Equal(t, "https://aidanhonan.com", rr.Result().Header.Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "Origin", rr.Result().Header.Get("Vary"))
	assert.Equal(t, "true", rr.Result().Header.Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With", rr.Result().Header.Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "POST", rr.Result().Header.Get("Access-Control-Allow-Methods"))

	// Assert API response
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `"test response"`, rr.Body.String())
}
