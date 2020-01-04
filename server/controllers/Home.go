package controllers

import (
	"Blog/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	//每次用户请求都去获取是否登录
	username := c.Query("username")
	//登出后，修改token
	token := c.Query("token")

	//先检查数据库是否存在token数据
	err := models.QueryTokenWightCon(username, token)
	if err > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "用户已登录", "username": username, "token": token})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "massage": "用户未登录"}) //不存在用户token
	}
}
