package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qqchat/common"
	"qqchat/model"
	"qqchat/models"
)

type UserBasic struct{}

// @Tags 用户基础信息
// @Summary 创建用户
// @Schemes
// @Description 创建用户基础信息
// @Accept json
// @Produce json
// @Router /api/users/user_basic/createUser [post]
// @Param data body model.CreateUserRequest true "请求参数"
// @Success 200 {object} common.Response
func (ub *UserBasic) CreateUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := common.ValidateRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.UserBasicModel.CreateUser(&req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"message": "创建用户成功",
	})
}

// @Tags 用户基础信息
// @Summary 用户列表
// @Schemes
// @Description 用户列表说明
// @Accept json
// @Produce json
// @Router /api/users/user_basic/getUsersList [get]
// @Param data body model.GetUsersListRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ub *UserBasic) GetUsersList(c *gin.Context) {
	var req model.GetUsersListRequest
	if err := common.ValidateRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err, users := models.UserBasicModel.GetUsersList(&req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, users)
}

// @Tags 用户基础信息
// @Summary 获取用户信息详情
// @Schemes
// @Description 获取用户信息详情说明
// @Accept json
// @Produce json
// @Router /api/users/user_basic/getUserInfo [get]
// @Param data body model.UserIdRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ub *UserBasic) GetUsersInfo(c *gin.Context) {
	var req model.UserIdRequest
	if err := common.ValidateRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err, userInfo := models.UserBasicModel.GetUsersInfo(&req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, userInfo)
}

// @Tags 用户基础信息
// @Summary 更新用户信息
// @Schemes
// @Description 更新用户信息说明
// @Accept json
// @Produce json
// @Router /api/users/user_basic/updateUser [post]
// @Param data body model.UpdateUserRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ub *UserBasic) UpdateUser(c *gin.Context) {
	var req model.UpdateUserRequest
	if err := common.ValidateRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.UserBasicModel.UpdateUser(&req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"message": "创建用户成功",
	})
}

// @Tags 用户基础信息
// @Summary 删除用户
// @Schemes
// @Description 删除用户接口说明
// @Accept json
// @Produce json
// @Router /api/users/user_basic/deleteUser [get]
// @Param data body model.UserIdRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ub *UserBasic) DeleteUser(c *gin.Context) {
	var req model.UserIdRequest
	if err := common.ValidateRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.UserBasicModel.DeleteUser(&req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"message": "删除用户成功",
	})
}
