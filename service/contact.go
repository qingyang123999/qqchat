package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qqchat/common"
	"qqchat/model"
	"qqchat/models"
)

type Contact struct{}

// @Tags 用户关系
// @Summary 创建用户关系
// @Schemes
// @Description 创建用户关系描述
// @Accept json
// @Produce json
// @Router /api/users/contact/createContact [post]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.CreateContactRequest true "请求参数"
// @Success 200 {object} common.Response
func (ct *Contact) CreateContact(c *gin.Context) {
	var req model.CreateContactRequest
	if err := common.ValidateRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.ContactModel.CreateContact(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"message": "创建用户关联关系成功",
	})
}

// @Tags 用户关系
// @Summary 用户关系列表
// @Schemes
// @Description 用户关系列表说明
// @Accept json
// @Produce json
// @Router /api/users/contact/getContactList [get]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.GetContactListRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ct *Contact) GetContactList(c *gin.Context) {
	var req model.GetContactListRequest
	if err := common.ValidateRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err, users := models.ContactModel.GetContactList(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, users)
}

// @Tags 用户关系
// @Summary 获取用户关系详情
// @Schemes
// @Description 获取用户关系详情详情说明
// @Accept json
// @Produce json
// @Router /api/users/contact/getContactInfo [get]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.ContactIdRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ct *Contact) GetContactInfo(c *gin.Context) {
	var req model.ContactIdRequest
	if err := common.ValidateRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//u, _ :=common.GetUserFromContext(c)
	//id:=u.ID
	err, userInfo := models.ContactModel.GetContactInfo(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, userInfo)
}

// @Tags 用户关系
// @Summary 更新用户关系
// @Schemes
// @Description 更新用户关系信息说明
// @Accept json
// @Produce json
// @Router /api/users/contact/updateContact [post]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.UpdateContactRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ct *Contact) UpdateContact(c *gin.Context) {
	var req model.UpdateContactRequest
	if err := common.ValidateRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.ContactModel.UpdateContact(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"message": "创建用户成功",
	})
}

// @Tags 用户关系
// @Summary 删除用户关系
// @Schemes
// @Description 删除用户关系接口说明
// @Accept json
// @Produce json
// @Router /api/users/contact/deleteContact [get]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.ContactIdRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ct *Contact) DeleteContact(c *gin.Context) {
	var req model.ContactIdRequest
	if err := common.ValidateRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.ContactModel.DeleteContact(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"message": "删除用户成功",
	})
}
