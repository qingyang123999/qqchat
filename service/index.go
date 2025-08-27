package service

import (
	"github.com/gin-gonic/gin"
	"qqchat/models"
	"qqchat/utils"
	"time"
)

func GetIndex(ctx *gin.Context) {

	model := models.UserBasic{
		Username:      "tan",
		Password:      "111",
		Phone:         "15874894579",
		Email:         "11435345@qq.com",
		Identity:      "dfdgdfgsd",
		ClientIp:      "125.22.00.123",
		ClientPort:    "80",
		LoginTime:     time.Now(),
		HeartbeatTime: time.Now(),
		LogoutTime:    time.Now(),
		IsLogout:      1,
		DeviceInfo:    "设备信息1",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err := utils.Db.Create(&model).Error
	if err != nil {
		err = nil
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get index",
	})
}
