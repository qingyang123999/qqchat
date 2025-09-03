package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"qqchat/common"
	"qqchat/model"
	"qqchat/utils"
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
func (ct *Contact) TableName() string {
	return "contact"
}

func (ct *Contact) CreateContact(c *gin.Context, req *model.CreateContactRequest) (err error) {
	// 创建数据库模型对象
	contactModel := &Contact{
		OwerId:   req.OwerId,
		TargetId: req.TargetId,
		Type:     req.Type,
		Desc:     req.Desc,
	}

	result := common.Db.Create(contactModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ct *Contact) GetContactList(c *gin.Context, req *model.GetContactListRequest) (err error, contacts []Contact) {
	// 构建查询条件，排除分页字段
	query := common.Db

	// 根据请求参数动态构建查询条件
	if req.OwerId > 0 {
		query = query.Where("ower_id = ?", req.OwerId)
	}
	if req.TargetId > 0 {
		query = query.Where("target_id = ?", req.TargetId)
	}
	if req.Type > 0 {
		query = query.Where("type = ?", req.Type)
	}

	// 执行查询
	result := query.Limit(req.PageSize).Offset(utils.GetPageOffset(req.Page, req.PageSize)).Find(&contacts)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, contacts
}

func (ct *Contact) GetContactInfo(c *gin.Context, req *model.ContactIdRequest) (err error, contactInfo Contact) {
	//context, err := common.GetUserFromContext(c)
	//if err != nil {
	//	return err, UserBasic{}
	//}
	//id:=context.ID
	result := common.Db.Where("id=?", req.ID).First(&contactInfo)
	if result.Error != nil {
		return result.Error, Contact{}
	}
	return nil, contactInfo
}

func (ct *Contact) UpdateContact(c *gin.Context, req *model.UpdateContactRequest) (err error) {
	// 将字符串时间转换为common.CustomTime
	updates := map[string]interface{}{
		"ower_id":   req.OwerId,
		"target_id": req.TargetId,
		"type":      req.Type,
		"desc":      req.Desc,
	}
	result := common.Db.Model(&Contact{}).Where("id = ?", req.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ct *Contact) DeleteContact(c *gin.Context, req *model.ContactIdRequest) (err error) {
	// 使用 Unscoped() 真正从数据库中删除记录，而不是软删除   方法二：保留 gorm.Model 但使用 Unscoped() 进行所有操作
	//或者 方法一：移除 gorm.Model 并手动定义字段（推荐 把数据表中的delete_at字段去掉
	result := common.Db.Unscoped().Where("id = ?", req.ID).Delete(&Contact{})
	if result.Error != nil {
		return result.Error
	}

	// 检查是否有记录被删除
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到ID为%d的用户记录", req.ID)
	}

	return nil
}
