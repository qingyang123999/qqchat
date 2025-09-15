package model

import "qqchat/utils"

type CreateMessagesRequest struct {
	ID       uint64 `gorm:"primaryKey;column:id"  json:"id"         form:"id"         binding:"required"`
	FormId   uint   `gorm:"column:form_id"        json:"formId"     form:"formId"     binding:"required"`
	TargetId uint   `gorm:"column:target_id"      json:"targetId"   form:"targetId"   binding:"required"`
	Type     uint   `gorm:"column:type"           json:"type"       form:"type"       binding:"required"`
	Media    string `gorm:"column:media"          json:"media"      form:"media"      binding:"required"`
	Content  string `gorm:"column:content"        json:"content"    form:"content"    binding:"required"`
	Pic      string `gorm:"column:pic"            json:"pic"        form:"pic"        binding:"omitempty"`
	Url      string `gorm:"column:url"            json:"url"        form:"url"        binding:"omitempty"`
	Desc     string `gorm:"column:desc"           json:"desc"       form:"desc"       binding:"omitempty"`
	Amount   uint64 `gorm:"column:amount"         json:"amount"     form:"amount"     binding:"omitempty"`
}

type GetMessagesListRequest struct {
	ID        uint64           `gorm:"primaryKey;column:id"  json:"id"         form:"id"         binding:"omitempty"`
	FormId    uint             `gorm:"column:form_id"        json:"formId"     form:"formId"     binding:"omitempty"`
	TargetId  uint             `gorm:"column:target_id"      json:"targetId"   form:"targetId"   binding:"omitempty"`
	Type      uint             `gorm:"column:type"           json:"type"       form:"type"       binding:"omitempty"`
	Media     string           `gorm:"column:media"          json:"media"      form:"media"      binding:"omitempty"`
	Content   string           `gorm:"column:content"        json:"content"    form:"content"    binding:"omitempty"`
	Pic       string           `gorm:"column:pic"            json:"pic"        form:"pic"        binding:"omitempty"`
	Url       string           `gorm:"column:url"            json:"url"        form:"url"        binding:"omitempty"`
	Desc      string           `gorm:"column:desc"           json:"desc"       form:"desc"       binding:"omitempty"`
	Amount    uint64           `gorm:"column:amount"         json:"amount"     form:"amount"     binding:"omitempty"`
	CreatedAt utils.CustomTime `gorm:"column:created_at"     json:"createdAtt" form:"createdAtt" binding:"omitempty"`
	UpdatedAt utils.CustomTime `gorm:"column:updated_at"     json:"updatedAt"  form:"updatedAt"  binding:"omitempty"`
	PageModel
}
type MessagesIdRequest struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
}

type SendMessagesRequest struct {
	FormId   uint `gorm:"column:form_id"        json:"formId"     form:"formId"     binding:"required"`
	TargetId uint `gorm:"column:target_id"      json:"targetId"   form:"targetId"   binding:"omitempty"`
}
