package service

var ApiService = apiService{
	UserBasic: UserBasic{},
}

type apiService = struct {
	UserBasic UserBasic
}
