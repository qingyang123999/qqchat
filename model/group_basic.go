package model

import "qqchat/utils"

type CreateGroupBasicRequest struct {
	Name   string `gorm:"column:name"           json:"name"       form:"name"      binding:"required"`
	OwerId uint   `gorm:"column:ower_id"        json:"owerId"     form:"owerId"    binding:"required"`
	Icon   string `gorm:"column:icon"           json:"icon"       form:"icon"      binding:"required"`
	Type   uint   `gorm:"column:type"           json:"type"       form:"type"      binding:"required"`
	Desc   string `gorm:"column:desc"           json:"desc"       form:"desc"      binding:"omitempty"`
}

type GetGroupBasicListRequest struct {
	ID        uint             `gorm:"primaryKey;column:id"  json:"id"         form:"id"        binding:"omitempty"`
	Name      string           `gorm:"column:name"           json:"name"       form:"name"      binding:"omitempty"`
	OwerId    uint             `gorm:"column:ower_id"        json:"owerId"     form:"owerId"    binding:"omitempty"`
	Icon      string           `gorm:"column:icon"           json:"icon"       form:"icon"      binding:"omitempty"`
	Type      uint             `gorm:"column:type"           json:"type"       form:"type"      binding:"omitempty"`
	Desc      string           `gorm:"column:desc"           json:"desc"       form:"desc"      binding:"omitempty"`
	CreatedAt utils.CustomTime `gorm:"column:created_at"     json:"createdAtt" form:"createdAt" binding:"omitempty"`
	UpdatedAt utils.CustomTime `gorm:"column:updated_at"     json:"updatedAt"  form:"updatedAt" binding:"omitempty"`
	PageModel
}
type GroupBasicIdRequest struct {
	ID uint `json:"id" form:"id" binding:"required"`
}
type UpdateGroupBasicRequest struct {
	ID        uint             `gorm:"primaryKey;column:id"  json:"id"         form:"id"        binding:"required"`
	Name      string           `gorm:"column:name"           json:"name"       form:"name"      binding:"omitempty"`
	OwerId    uint             `gorm:"column:ower_id"        json:"owerId"     form:"owerId"    binding:"omitempty"`
	Icon      string           `gorm:"column:icon"           json:"icon"       form:"icon"      binding:"omitempty"`
	Type      uint             `gorm:"column:type"           json:"type"       form:"type"      binding:"omitempty"`
	Desc      string           `gorm:"column:desc"           json:"desc"       form:"desc"      binding:"omitempty"`
	CreatedAt utils.CustomTime `gorm:"column:created_at"     json:"createdAtt" form:"createdAt" binding:"omitempty"`
	UpdatedAt utils.CustomTime `gorm:"column:updated_at"     json:"updatedAt"  form:"updatedAt" binding:"omitempty"`
}
