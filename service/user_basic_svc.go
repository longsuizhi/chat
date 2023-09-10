package service

import (
	"chat/api/code"
	"chat/dao"
	"chat/logger"
	"chat/models"
	"chat/utils"
	"fmt"
	"math/rand"
	"time"

	"github.com/asaskevich/govalidator"
	"go.uber.org/zap"
)

func GetUserList() ([]dao.UserBasic, error){
	var err error
	res, err := dao.GetUserList()
	if err != nil {
		logger.Logger.Error("dao.GetUserList failed", zap.Error(err))
	}
	return res, err
}

func CreateUser(req models.CreateUserReq) (dao.UserBasic, error){
	res := dao.UserBasic{}
	//user.Name = c.Query("name")
	//password := c.Query("password")
	//repassword := c.Query("repassword")
	salt := fmt.Sprintf("%06d",rand.Int31())
	data, err  := dao.FindUserByName(req.Name)
	if err != nil {
		logger.Logger.Error("dao.FindUserByName failed", zap.Error(err))
	}
	//用户名已存在
	if data.ID > 0 {
		return res, code.UserNameExist
	}
	if req.PassWord != req.RePassWord {
		return res, code.PasswordInconsistency
	}
	//生成随机唯一标识
	rand.Seed(time.Now().UnixNano())
	var account string 
	for{
		chars := "0123456789"
		result := make([]byte, 10)
		for i := range result {
			result[i] = chars[rand.Intn(len(chars))]
		}
		account = string(result)
		user, err := dao.FindUserByIdentity(account)
		if err != nil {
			logger.Logger.Error("dao.FindUserByIdentity", zap.Error(err))
		}
		if user.ID <= 0 {
			break
		}
	}
	newUser := dao.UserBasic{
		Name: req.Name,
		PassWord: utils.MakePassword(req.PassWord, salt),
		Salt: salt,
		Identity: account,
		LoginTime: time.Now(),
		LoginOutTime: time.Now(),
		HeartbeatTime: time.Now(),
	}
	createErr := dao.CreateUser(&newUser)
	if createErr != nil {
		logger.Logger.Error("dao.CreateUser failed", zap.Error(createErr))
	}
	return res, createErr
}

func UpdateUser(req models.UpdateUserReq) (dao.UserBasic, error){
	res := dao.UserBasic{}
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		logger.Logger.Error("govalidator.ValidateStruct", code.ParametersDoNotMatch)
		return res, err
	}
	user, err := dao.FindUserByIdentity(req.Identity)
	if err != nil {
		logger.Logger.Error("dao.FindUserByIdentity", zap.Error(err))
		return res, err
	}
	if user.ID <= 0 {
		return res, code.UserNotExist
	}
	user, err = dao.FindUserByName(req.Name)
	if err != nil {
		logger.Logger.Error("dao.FindUserByName", zap.Error(err))
		return res, err
	}
	//用户名已存在并且不是该账号的用户名
	if user.ID > 0 || user.Identity != req.Identity {
		return res, code.UserNameExist
	}
	user.Name = req.Name
	user.Phone = req.Phone
	user.Email = req.Email
	user.Avatar = req.Avatar
	updateErr := dao.UpdateUser(&user)
	if updateErr != nil {
		logger.Logger.Error("dao.UpdateUser", zap.Error(updateErr))

	}
	return user, err
}