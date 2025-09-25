package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"qqchat/common"
	"qqchat/model"
	"qqchat/models"
	"time"
)

type UserBasic struct{}

// @Tags 用户基础信息
// @Summary 创建用户
// @Schemes
// @Description 创建用户基础信息
// @Accept json
// @Produce json
// @Router /api/users/user_basic/createUser [post]
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.CreateUserRequest true "请求参数"
// @Success 200 {object} common.Response
func (ub *UserBasic) CreateUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := common.ValidateJSONRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.UserBasicModel.CreateUser(c, &req)
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
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.GetUsersListRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ub *UserBasic) GetUsersList(c *gin.Context) {
	var req model.GetUsersListRequest
	if err := common.ValidateQueryRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err, users := models.UserBasicModel.GetUsersList(c, &req)
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
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.UserIdRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ub *UserBasic) GetUsersInfo(c *gin.Context) {
	var req model.UserIdRequest
	if err := common.ValidateQueryRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//u, _ :=common.GetUserFromContext(c)
	//id:=u.ID
	err, userInfo := models.UserBasicModel.GetUsersInfo(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userInfo.Password = ""
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
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.UpdateUserRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ub *UserBasic) UpdateUser(c *gin.Context) {
	var req model.UpdateUserRequest
	if err := common.ValidateJSONRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.UserBasicModel.UpdateUser(c, &req)
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
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.UserIdRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ub *UserBasic) DeleteUser(c *gin.Context) {
	var req model.UserIdRequest
	if err := common.ValidateQueryRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.UserBasicModel.DeleteUser(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"message": "删除用户成功",
	})
}

// @Tags 用户基础信息
// @Summary 用户登录
// @Schemes
// @Description 用户登录接口说明
// @Accept json
// @Produce json
// @Router /api/users/user_basic/login [post]
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.LoginRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ub *UserBasic) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := common.ValidateJSONRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err, token := models.UserBasicModel.Login(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"token": token,
	})
}

// @Tags 用户基础信息
// @Summary 用户登录并获取用户信息
// @Schemes
// @Description 用户登录并获取用户信息接口说明
// @Accept json
// @Produce json
// @Router /api/users/user_basic/findUserByNameAndPwd [post]
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.LoginRequest1 true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (ub *UserBasic) FindUserByNameAndPwd(c *gin.Context) {
	var req model.LoginRequest1
	if err := common.ValidateRequest(c, &req); err != nil {
		common.ErrorResponse(c, -1, err.Error())
		return
	}

	err, token, userInfo := models.UserBasicModel.FindUserByNameAndPwd(c, &req)
	if err != nil {
		common.ErrorResponse(c, -1, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, struct {
		Identity string `json:"Identity"`
		ID       uint   `json:"Id"`
		models.UserBasic
	}{
		Identity:  token,
		ID:        userInfo.ID,
		UserBasic: userInfo,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {

	// user.Name = c.Query("name")
	// password := c.Query("password")
	// repassword := c.Query("repassword")
	user := models.UserBasic{}
	user.Username = c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("Identity")
	salt := fmt.Sprintf("%06d", rand.Int31())

	err, data := models.UserBasicModel.GetUsersInfoByUserName(user.Username)
	if err != nil {
		common.ErrorResponse(c, -1, err.Error())
	}
	if user.Username == "" || password == "" || repassword == "" {
		common.ErrorResponse(c, -1, "用户名或密码不能为空！")
		return
	}
	if data.Username != "" {
		common.ErrorResponse(c, -1, "用户名已注册！")
		return
	}
	if password != repassword {
		common.ErrorResponse(c, -1, "两次密码不一致！")
		return
	}
	//user.PassWord = password
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	fmt.Println(user.PassWord)
	user.LoginTime = time.Now()
	user.LoginOutTime = time.Now()
	user.HeartbeatTime = time.Now()
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "新增用户成功！",
		"data":    user,
	})
}
