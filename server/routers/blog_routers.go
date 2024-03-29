/*
	路由器部分
*/

package routers

import (
	"Blog/server/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())

	{
		router.POST("/registerPost", controllers.RegisterPost)   // 注册
		router.GET("/exit", controllers.ExitGet)                 //退出
		router.POST("/", controllers.Home)                       // 首页
		router.POST("/loginpost", controllers.LoginPost)         //登录
		router.POST("/createArticle", controllers.CreateArticle) //创建文章
		router.POST("/deletearticle", controllers.DeleteArticle) //删除文章
	}

	// v1 := router.Group("/article")
	// {
	// 	v1.POST("/insertarticle", controllers.insertarticle)
	// }

	return router

}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
