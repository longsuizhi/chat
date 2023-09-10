package dao

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	Model
	Name          string    `json:"name"`
	PassWord      string    `json:"password"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Avatar        string    `json:"avatar"`   //头像
	Identity      string    `json:"identity"` //唯一标识
	ClientIp      string    `json:"client_ip"`
	ClientPort    string    `json:"client_port"`
	Salt          string    `json:"salt"` //加密用的随机数
	LoginTime     time.Time `json:"login_time"`
	HeartbeatTime time.Time `json:"heartbeat_time"`
	LoginOutTime  time.Time `json:"login_out_time"`
	IsLogout      bool      `json:"is_logout"`
	DeviceInfo    string    `json:"device_info"` //设备信息
}

func GetUserList() ([]UserBasic, error) {
	res := []UserBasic{}
	err := DB.Model(UserBasic{}).Find(&res).Limit(10).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, nil
	}
	return res, err
}

func FindUserByName(name string) (UserBasic, error) {
	res := UserBasic{}
	err := DB.Table("user_basics").Where("name = ? and deleted_at is null", name).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, nil
	}
	return res, err
}

func FindUserByIdentity(account string) (UserBasic, error) {
	res := UserBasic{}
	err := DB.Model(UserBasic{}).Where("identity = ?", account).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, nil
	}
	return res, err
}

func CreateUser(newUser *UserBasic) error {
	return DB.Model(UserBasic{}).Create(newUser).Error
}

func UpdateUser(user *UserBasic) error {
	return DB.Model(UserBasic{}).Where("id = ?", user.ID).Updates(user).Error
}
