package main

import (
	"Blog/server/database"
	"Blog/server/routers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.Initmysql()
	router := routers.InitRouter()

	router.Run()
}
