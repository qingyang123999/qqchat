package service

// 注册所有的 控制器的所有结构体
var ApiService = apiService{
	UserBasic:    UserBasic{},
	Contact:      Contact{},
	GroupBasic:   GroupBasic{},
	Messages:     Messages{},
	SysWebSocket: SysWebSocket{},
}

type apiService = struct {
	UserBasic    UserBasic
	Contact      Contact
	GroupBasic   GroupBasic
	Messages     Messages
	SysWebSocket SysWebSocket
}
