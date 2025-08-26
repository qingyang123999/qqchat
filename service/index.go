package service

import (
	"github.com/gin-gonic/gin"
	"qqchat/models"
	"qqchat/utils"
)

func GetIndex(ctx *gin.Context) {

	model := models.UserBasic{
		Id: 2,
	}
	err := utils.Db.Take(&model, "id = ?", 3).Error
	if err == nil {
		return
	}

	err = utils.Db.Create(&model).Error
	if err != nil {
		err = nil
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get index",
	})
}
