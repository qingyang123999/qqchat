package model

type LoginRequest struct {
	Username string `json:"username" form:"username"  binding:"required,min=2,max=20"`
	Password string `json:"password" form:"password" binding:"required,min=8"`
}

type LoginRequest1 struct {
	Name     string `json:"name"       form:"name"  binding:"required,min=2,max=20"`
	Password string `json:"password"   form:"password"  binding:"required,min=8"`
}

type CreateUserRequest struct {
	Username      string `json:"username"       form:"username"      binding:"required,min=2,max=20,alphanum"`
	Password      string `json:"password"       form:"password"      binding:"required,min=8,max=64,containsany=!@#$%^&*"`
	Phone         string `json:"phone"          form:"phone"         binding:"required,mobile"`
	Email         string `json:"email"          form:"email"         binding:"required,email"              `
	Identity      string `json:"identity"       form:"identity"      binding:"omitempty"                   `
	ClientIp      string `json:"clientIp"       form:"clientIp"      binding:"omitempty,ip"                `
	ClientPort    string `json:"clientPort"     form:"clientPort"    binding:"omitempty,numeric"           `
	LoginTime     string `json:"loginTime"      form:"loginTime"     binding:"omitempty,datetime=2006-01-02 15:04:05"`
	HeartbeatTime string `json:"heartbeatTime"  form:"heartbeatTime" binding:"omitempty,datetime=2006-01-02 15:04:05"`
	LogoutTime    string `json:"logoutTime"     form:"logoutTime"    binding:"omitempty,datetime=2006-01-02 15:04:05"`
	IsLogout      int    `json:"isLogout"       form:"isLogout"      binding:"omitempty,oneof=0 1"         `
	DeviceInfo    string `json:"deviceInfo"     form:"deviceInfo"    binding:"omitempty,max=256"           `
}

type GetUsersListRequest struct {
	Username      string `json:"username"       form:"username"         binding:"omitempty,min=2,max=20,alphanum"`
	Phone         string `json:"phone"          form:"phone"         binding:"omitempty,mobile"`
	Email         string `json:"email"          form:"email"         binding:"omitempty,email"             `
	Identity      string `json:"identity"       form:"identity"      binding:"omitempty"                   `
	ClientIp      string `json:"clientIp"       form:"clientIp"      binding:"omitempty,ip"                `
	ClientPort    string `json:"clientPort"     form:"clientPort"    binding:"omitempty,numeric"           `
	LoginTime     string `json:"loginTime"      form:"loginTime"      binding:"omitempty,datetime=2006-01-02 15:04:05"`
	HeartbeatTime string `json:"heartbeatTime"  form:"heartbeatTime"  binding:"omitempty,datetime=2006-01-02 15:04:05"`
	LogoutTime    string `json:"logoutTime"     form:"logoutTime"     binding:"omitempty,datetime=2006-01-02 15:04:05"`
	IsLogout      int    `json:"isLogout"       form:"isLogout"      binding:"omitempty,oneof=0 1"         `
	DeviceInfo    string `json:"deviceInfo"     form:"deviceInfo"    binding:"omitempty,max=256"           `
	PageModel
}
type UserIdRequest struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
}
type UpdateUserRequest struct {
	ID            uint64 `json:"id"             form:"id"            binding:"required"`
	Username      string `json:"username"       form:"username"      binding:"omitempty,min=2,max=20,alphanum"`
	Password      string `json:"password"       form:"password"      binding:"omitempty,min=8,max=64,containsany=!@#$%^&*"`
	Phone         string `json:"phone"          form:"phone"         binding:"omitempty,mobile"`
	Email         string `json:"email"          form:"email"         binding:"omitempty,email"              `
	Identity      string `json:"identity"       form:"identity"      binding:"omitempty"                   `
	ClientIp      string `json:"clientIp"       form:"clientIp"      binding:"omitempty,ip"                `
	ClientPort    string `json:"clientPort"     form:"clientPort"    binding:"omitempty,numeric"           `
	LoginTime     string `json:"loginTime"      form:"loginTime"      binding:"omitempty,datetime=2006-01-02 15:04:05"`
	HeartbeatTime string `json:"heartbeatTime"  form:"heartbeatTime"  binding:"omitempty,datetime=2006-01-02 15:04:05"`
	LogoutTime    string `json:"logoutTime"     form:"logoutTime"     binding:"omitempty,datetime=2006-01-02 15:04:05"`
	IsLogout      int    `json:"isLogout"       form:"isLogout"      binding:"omitempty,oneof=0 1"         `
	DeviceInfo    string `json:"deviceInfo"     form:"deviceInfo"    binding:"omitempty,max=256"           `
}
