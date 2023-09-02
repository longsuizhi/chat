package models

import (
	"chat/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Avatar        string //头像
	Identity      string	//唯一标识
	ClientIp      string
	ClientPort    string
	Salt          string	//加密用的随机数
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time
	IsLogout      bool
	DeviceInfo    string //设备信息
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

// 获取用户列表
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

// 通过name查找用户
func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

// 创建用户
func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}