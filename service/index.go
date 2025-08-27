package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qqchat/models"
	"qqchat/utils"
	"time"
)

// @Tags 目录名称
// @Summary 接口名称
// @Schemes
// @Description 接口说明
// @Accept json
// @Produce json
// @Router /api/v1/example/helloworld [get]
// @Param data body models.AAAA true "请求参数  @Param [参数名] [参数类型] [数据类型] [是否必填] [描述信息(可选)]"
// @Success 200 {string} helloworld
// @Failure 500  {string} helloworld
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @Tags 测试目录
// @Summary 测试接口
// @Schemes
// @Description 用于测试
// @Accept json
// @Produce json
// @Router /index [get]
// @Param data body models.AAAA true "请求参数"
// @Success   200  {object}  models.AAAA
// @Failure   500  {object}  models.AAAA
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
