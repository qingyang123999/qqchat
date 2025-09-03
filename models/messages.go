package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"qqchat/common"
	"qqchat/model"
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

func (m *Messages) CreateMessages(c *gin.Context, req *model.CreateMessagesRequest) (err error) {
	// 创建数据库模型对象
	messagesModel := &Messages{
		FormId:   req.FormId,
		TargetId: req.TargetId,
		Type:     req.Type,
		Media:    req.Media,
		Content:  req.Content,
		Pic:      req.Pic,
		Url:      req.Url,
		Desc:     req.Desc,
		Amount:   req.Amount,
	}

	result := common.Db.Create(messagesModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *Messages) GetMessagesList(c *gin.Context, req *model.GetMessagesListRequest) (err error, messages []Messages) {
	// 构建查询条件，排除分页字段
	query := common.Db

	// 根据请求参数动态构建查询条件
	if req.FormId > 0 {
		query = query.Where("form_id = ?", req.FormId)
	}
	if req.TargetId > 0 {
		query = query.Where("target_id = ?", req.TargetId)
	}
	if req.Type > 0 {
		query = query.Where("type = ?", req.Type)
	}
	if req.Media != "" {
		query = query.Where("media = ?", req.Media)
	}
	if req.Content != "" {
		query = query.Where("content = ?", req.Content)
	}
	if req.Amount > 0 {
		query = query.Where("amount = ?", req.Amount)
	}

	// 执行查询
	result := query.Limit(req.PageSize).Offset(utils.GetPageOffset(req.Page, req.PageSize)).Find(&messages)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, messages
}

func (m *Messages) GetMessagesInfo(c *gin.Context, req *model.MessagesIdRequest) (err error, messagesInfo Messages) {
	//context, err := common.GetMessagesFromContext(c)
	//if err != nil {
	//	return err, MessagesBasic{}
	//}
	//id:=context.ID
	result := common.Db.Where("id=?", req.ID).First(&messagesInfo)
	if result.Error != nil {
		return result.Error, Messages{}
	}
	return nil, messagesInfo
}

func (m *Messages) DeleteMessages(c *gin.Context, req *model.MessagesIdRequest) (err error) {
	// 使用Model指定模型，然后通过主键ID删除
	result := common.Db.Model(&Messages{}).Where("id = ?", req.ID).Delete(&Messages{})
	if result.Error != nil {
		return result.Error
	}

	// 检查是否有记录被删除
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到ID为%d的记录", req.ID)
	}

	return nil
}
