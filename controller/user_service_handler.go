package controller

import (
	"chat/api/code"
	"chat/logger"
	"chat/models"
	"chat/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [post]
func GetUserListHandler(c *gin.Context) {
	data, err := service.GetUserList()
	if err != nil {
		logger.Logger.Error("service.GetUserList failed", zap.Error(err))
	}
	code.Response(c, err, data)
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [post]
func CreateUserHandler(c *gin.Context) {
	req := models.CreateUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Error("UserRegisterHandler with invalid param", zap.Error(err))
		code.Response(c, err, nil)
		return
	}
	if req.Name == "" || req.PassWord == "" || req.RePassWord == "" {
		logger.Logger.Error("CreateUserHandler with invalid param", zap.Error(code.UserPasswordIsNull))
		code.Response(c, code.UserPasswordIsNull, nil)
		return
	}
	data, err := service.CreateUser(req)
	if err != nil {
		logger.Logger.Error("service.CreateUser failed", zap.Error(err))
	}
	code.Response(c, err, data)
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUserHandler(c *gin.Context) {
	req := models.UpdateUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil || req.Identity == "" || req.Name == "" || req.Phone == "" || req.Email == "" || req.Avatar == ""{
		logger.Logger.Error("UpdateUserHandler with invalid param", code.InvalidParam)
		code.Response(c, code.InvalidParam, nil)
		return
	}
	data, err := service.UpdateUser(req)
	if err != nil {
		logger.Logger.Error("service.UpdateUser failed", zap.Error(err))
	}
	code.Response(c, err, data)
}