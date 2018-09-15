package middleware_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/karuppaiah/shopping/middleware"
	"github.com/stretchr/testify/assert"
)

func GetTestRouter() *gin.Engine {
	r := gin.Default()

	return r
}

func TestCORS(t *testing.T) {
	r := GetTestRouter()
	m := middleware.InitMiddleware()
	r.Use(m.CORS())
	rec := httptest.NewRecorder()
	r.GET("/", func(c *gin.Context) {
	})

	h := m.CORS()
	c, _ := gin.CreateTestContext(rec)

	h(c)
	assert.Equal(t, "*", rec.Header().Get("Access-Control-Allow-Origin"))
}

// func TestAuth(t *testing.T) {

// 	m := middleware.InitMiddleware()

// 	rec := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(rec)

// 	c.Request = httptest.NewRequest("GET", "/product", nil)
// 	h := m.AuthMiddleware()
// 	call := h.Authenticator
// 	res, _ := call(c)

// 	assert.Equal(t, nil, nil)
// }
