package models

import (
	"gorm.io/gorm"
	"qqchat/utils"
)

var MessagesModel = Messages{}

type Messages struct {
	gorm.Model
	ID        uint64           `gorm:"primaryKey;column:id"  json:"id"`
	FormId    uint             `gorm:"column:form_id"        json:"formId"`
	TargetId  uint             `gorm:"column:target_id"      json:"targetId"`
	Type      uint             `gorm:"column:type"           json:"type"`
	Media     string           `gorm:"column:media"          json:"media"`
	Content   string           `gorm:"column:content"        json:"content"`
	Pic       string           `gorm:"column:pic"            json:"pic"`
	Url       string           `gorm:"column:url"            json:"url"`
	Desc      string           `gorm:"column:desc"           json:"desc"`
	Amount    uint64           `gorm:"column:amount"         json:"amount"`
	CreatedAt utils.CustomTime `gorm:"column:created_at"     json:"createdAtt"`
	UpdatedAt utils.CustomTime `gorm:"column:updated_at"     json:"updatedAt"`
}

// 设置表名称  默认的表明会带s  链接表名会变成users
func (m *Messages) TableName() string {
	return "messages"
}
