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

	{
		router.POST("/registerPost", controllers.RegisterPost)
		router.GET("/exit", controllers.ExitGet)
		router.GET("/", controllers.Home)
		router.POST("/loginpost", controllers.LoginPost)
	}

	// v1 := router.Group("/article")
	// {
	// 	v1.POST("/insertarticle", controllers.insertarticle)
	// }

	return router

}
