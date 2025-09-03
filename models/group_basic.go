package models

import (
	"gorm.io/gorm"
	"qqchat/utils"
)

var GroupBasicModel = GroupBasic{}

type GroupBasic struct {
	gorm.Model
	ID        uint64           `gorm:"primaryKey;column:id"  json:"id"`
	Name      string           `gorm:"column:name"           json:"name"`
	OwerId    uint             `gorm:"column:ower_id"        json:"owerId"`
	Icon      string           `gorm:"column:icon"           json:"icon"`
	Type      uint             `gorm:"column:type"           json:"type"`
	Desc      string           `gorm:"column:desc"           json:"desc"`
	Amount    uint64           `gorm:"column:amount"         json:"amount"`
	CreatedAt utils.CustomTime `gorm:"column:created_at"     json:"createdAtt"`
	UpdatedAt utils.CustomTime `gorm:"column:updated_at"     json:"updatedAt"`
}

// 设置表名称  默认的表明会带s  链接表名会变成users
func (g *GroupBasic) TableName() string {
	return "group_basic"
}
