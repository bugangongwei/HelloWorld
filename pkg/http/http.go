package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitHttpEngine init a gin Engine and run it
func InitHttpEngine() *gin.Engine {
	engine := gin.New()

	// middleware
	engine.Use(LogRW)

	// 注册 handler
	engine.Handle(http.MethodGet, "/test", GetHistories)

	engine.Run("127.0.0.1:8080")

	return engine
}
