package service

import (
	"chat/models"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([] *models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}