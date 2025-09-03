package models

import (
	"gorm.io/gorm"
)

var ContactModel = Contact{}

type Contact struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;column:id"  json:"id"`
	OwerId   uint   `gorm:"column:ower_id"        json:"owerId"`
	TargetId uint   `gorm:"column:target_id"      json:"targetId"`
	Type     uint   `gorm:"column:type"           json:"type"`
	Desc     string `gorm:"column:desc"           json:"desc"`
}

// 设置表名称  默认的表明会带s  链接表名会变成users
func (c *Contact) TableName() string {
	return "contact"
}
