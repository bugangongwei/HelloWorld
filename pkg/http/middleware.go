package http

import (
	"github.com/gin-gonic/gin"
)

func LogRW(c *gin.Context) {
	// log.Infof("start to handle requests")
	c.Next()
	// log.Infof("end handle requests")
}
