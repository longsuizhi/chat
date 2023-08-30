package main

import (
	"chat/models"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
  db, err := gorm.Open(mysql.Open("user:123456@tcp(43.143.231.157:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // 迁移 schema
  db.AutoMigrate(&models.UserBasic{})

  // Create
  user := &models.UserBasic{}
  user.Name = "longsuizhi"
  user.LoginOutTime = time.Now()
  user.HeartbeatTime = time.Now()
  user.LoginTime = time.Now()
  db.Create(user)

  // Read
  fmt.Println(db.First(user, 1)) // 根据整型主键查找

  // Update - 将 product 的 price 更新为 200
  db.Model(&user).Update("PassWord", "1234")
}