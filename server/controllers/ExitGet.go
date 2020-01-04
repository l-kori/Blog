package controllers

import (
	"Blog/server/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ExitGet(c *gin.Context) {
	username := c.Query("username")
	//登出后，修改token
	token := c.Query("token")

	//先检查数据库是否存在token数据
	err := models.QueryTokenWightCon(username, token)
	//如果存在，则去删除token
	if err > 0 {
		err := models.DeleteTokenWithUsername(username, token)
		fmt.Println(err)

	} else {
		fmt.Println("用户已退出") //token不存在
	}

}
