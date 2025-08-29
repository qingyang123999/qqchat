package service

// 注册所有的 控制器的所有结构体
var ApiService = apiService{
	UserBasic: UserBasic{},
}

type apiService = struct {
	UserBasic UserBasic
}
