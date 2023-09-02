package service

import (
	"chat/models"
	"chat/utils"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [post]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"code": 0,
		"message": "用户名已注册！",
		"data": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	//user.Name = c.Query("name")
	//password := c.Query("password")
	//repassword := c.Query("repassword")
	user.Name = c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("repassword")
	fmt.Println(user.Name, " >>>>>>>>>> ", password, repassword)
	salt := fmt.Sprintf("%06d",rand.Int31())
	data := models.FindUserByName(user.Name)
	if user.Name == "" || password == "" || repassword == "" {
		c.JSON(200, gin.H{
			"code": -1,//  0成功   -1失败
			"message": "用户名或密码不能为空！",
			"data":    user,
		})
		return
	}
	if data.Name != "" {
		c.JSON(200, gin.H{
			"code": -1,
			"message": "用户名已注册！",
			"data": user,
		})
		return
	}
	if password != repassword {
		c.JSON(200, gin.H{
			"code": -1,
			"message": "两次密码不一致！",
			"data": user,
		})
		return
	}
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt;
	fmt.Print(user.PassWord)
	user.LoginTime = time.Now()
	user.LoginOutTime = time.Now()
	user.HeartbeatTime = time.Now()
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code" : 0,
		"message" : "新增用户成功！",
		"data": user,
	})
}