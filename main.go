package main

import (
	"chat/dao"
	"chat/logger"
	"chat/router"
)


func main() {
	dao.InitConfig()
	dao.InitMySQL()
	logger.Logger = logger.InitLogger()
	r := router.Router()
	r.Run(":8082")
}