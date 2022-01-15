package http

import (
	"git.llsapp.com/awesome/log"
	"github.com/gin-gonic/gin"
)

func LogRW(c *gin.Context) {
	log.Infof("start to handle requests")
	c.Next()
	log.Infof("end handle requests")
}
