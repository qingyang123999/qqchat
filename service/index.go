package service

import (
	"net/http"
	"qqchat/common"
	"qqchat/model"
	"qqchat/models"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	temp, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(c.Writer, "index")
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}

func ToRegister(c *gin.Context) {
	temp, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(c.Writer, "register")
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}

func ToChat(c *gin.Context) {
	temp, err := template.ParseFiles("views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/main.html")
	if err != nil {
		panic(err)
	}
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = token
	//fmt.Println("ToChat>>>>>>>>", user)
	temp.Execute(c.Writer, user)
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}

func Chat(c *gin.Context) {
	var req model.SendMessagesRequest
	if err := common.ValidateQueryRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	models.Chat(c.Writer, c.Request, req)
}
