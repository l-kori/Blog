package controllers

import (
	"Blog/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExitGet(c *gin.Context) {
	username := c.Query("username")
	//登出后，修改token
	token_id := c.Query("token")

	//先检查数据库是否存在token数据
	err := models.QueryTokenWithParam(username)
	//如果存在，则去删除token
	if err > 0 {
		err := models.DeleteTokenWithUsername(username, token_id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "massage": "退出失败"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 200, "massage": "退出成功"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "massage": "未登录"}) //token不存在
	}
}
