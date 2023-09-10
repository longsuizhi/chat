package models

import (
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
	Identity      string //唯一标识
	ClientIp      string
	ClientPort    string
	Salt          string //加密用的随机数
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time
	IsLogout      bool
	DeviceInfo    string //设备信息
}

type CreateUserReq struct {
	Name       string `json:"name"`
	PassWord   string `json:"password"`
	RePassWord string `json:"repassword"`
}

type UpdateUserReq struct {
	Identity string `json:"identity"` //唯一标识
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Avatar   string `json:"icon"` //头像
}
