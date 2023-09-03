package main

import (
	"chat/router"

	"chat/utils"
)


func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run(":8082")
}