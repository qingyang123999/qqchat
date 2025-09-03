package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qqchat/common"
	"qqchat/model"
	"qqchat/models"
)

type GroupBasic struct{}

// @Tags 群基础信息
// @Summary 创建群基础信息
// @Schemes
// @Description 创建群基础信息接口说明
// @Accept json
// @Produce json
// @Router /api/users/group_basic/createGroupBasic [post]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.CreateGroupBasicRequest  true "请求参数"
// @Success 200 {object} common.Response
func (gb *GroupBasic) CreateGroupBasic(c *gin.Context) {
	var req model.CreateGroupBasicRequest
	if err := common.ValidateJSONRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.GroupBasicModel.CreateGroupBasic(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"message": "创建成功",
	})
}

// @Tags 群基础信息
// @Summary 群基础信息列表
// @Schemes
// @Description 群基础信息说明
// @Accept json
// @Produce json
// @Router /api/users/group_basic/getGroupBasicsList [get]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.GetGroupBasicListRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (gb *GroupBasic) GetGroupBasicsList(c *gin.Context) {
	var req model.GetGroupBasicListRequest
	if err := common.ValidateQueryRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err, users := models.GroupBasicModel.GetGroupBasicsList(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, users)
}

// @Tags 群基础信息
// @Summary 获取群基础信息详情
// @Schemes
// @Description 获取群基础信息详情说明
// @Accept json
// @Produce json
// @Router /api/users/group_basic/getGroupBasicsInfo [get]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.GroupBasicIdRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (gb *GroupBasic) GetGroupBasicsInfo(c *gin.Context) {
	var req model.GroupBasicIdRequest
	if err := common.ValidateQueryRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//u, _ :=common.GetUserFromContext(c)
	//id:=u.ID
	err, userInfo := models.GroupBasicModel.GetGroupBasicsInfo(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, userInfo)
}

// @Tags 群基础信息
// @Summary 更新群基础信息
// @Schemes
// @Description 更新群基础信息说明
// @Accept json
// @Produce json
// @Router /api/users/group_basic/updateGroupBasic [post]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.UpdateGroupBasicRequest  true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (gb *GroupBasic) UpdateGroupBasic(c *gin.Context) {
	var req model.UpdateGroupBasicRequest
	if err := common.ValidateJSONRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.GroupBasicModel.UpdateGroupBasic(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"message": "创建用户成功",
	})
}

// @Tags 群基础信息
// @Summary 删除群基础信息
// @Schemes
// @Description 删除群基础信息接口说明
// @Accept json
// @Produce json
// @Router /api/users/group_basic/deleteGroupBasic [get]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.GroupBasicIdRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (gb *GroupBasic) DeleteGroupBasic(c *gin.Context) {
	var req model.GroupBasicIdRequest
	if err := common.ValidateQueryRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.GroupBasicModel.DeleteGroupBasic(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"message": "删除用户成功",
	})
}
