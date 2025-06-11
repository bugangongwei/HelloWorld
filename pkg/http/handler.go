package http

import (
	"net/http"

	"bugangongwei/HelloWorld/repo"
	"bugangongwei/HelloWorld/util"

	"github.com/gin-gonic/gin"
)

// GetHistories 打印 util.BindQuery 值, 用来调试 util 的方法
func GetHistories(c *gin.Context) {
	req := repo.GetHistoriesRequest{}
	if err := util.BindQuery(c, &req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// log.Infof("GetHistories request: %v", req)

	c.Data(http.StatusOK, "application/json", nil)
}
