package utils

import (
	"Blog/server/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func QueryToken(c *gin.Context, username string, token_id string) {
	err := models.QueryTokenWihtCon(username, token_id)
	if err > 0 {
		token_due_time := models.QueryToken_due_timeWithusername(username)
		now_token_due_time := time.Now().Unix()
		token_time_count := now_token_due_time - token_due_time
		//如果相差时间大于24小时 则删除token，并提示已退出
		if token_time_count > 50 { //86400
			err := models.DeleteTokenWithUsername(username, token_id)
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "用户已退出"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "用户已登录", "username": username, "token": token_id})

		}

	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "massage": "用户未登录"}) //不存在用户token
	}
}
