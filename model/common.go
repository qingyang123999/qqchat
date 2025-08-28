package model

type PageModel struct {
	Page     int `json:"page" form:"page" binding:"required,min=1"`
	PageSize int `json:"pageSize" form:"pageSize" binding:"required,min=1"`
}
