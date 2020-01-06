package controllers

import (
	"Blog/server/utils"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	//每次用户请求都去获取是否登录
	username := c.Query("username")
	//登出后，修改token
	token := c.Query("token")

	//先检查数据库是否存在token数据以及token存在时间
	utils.QueryToken(c, username, token)
}
