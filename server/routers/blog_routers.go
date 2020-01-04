/*
	路由器部分
*/

package routers

import (
	"Blog/server/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//store := cookie.NewStore([]byte("loginuser"))
	//router.Use(sessions.Sessions("mysession",store))

	{
		router.POST("/registerPost", controllers.RegisterPost)
		router.GET("/exit", controllers.ExitGet)
		router.GET("/", controllers.Home)
		//router.GET("/register",controllers.RegisterGet)
		////router.POST("/pwd",controllers.RegisterPost)
		router.POST("/loginpost", controllers.LoginPost)
	}

	//v1 := router.Group("/article")
	//{
	//	v1.GET("/add", controllers.AddArticleGet)
	//	v1.POST("/add", controllers.AddArticlePost)
	//}

	return router

}
