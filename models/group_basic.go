package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"qqchat/common"
	"qqchat/model"
	"qqchat/utils"
)

var GroupBasicModel = GroupBasic{}

type GroupBasic struct {
	gorm.Model
	ID        uint             `gorm:"primaryKey;column:id"  json:"id"`
	Name      string           `gorm:"column:name"           json:"name"`
	OwerId    uint             `gorm:"column:ower_id"        json:"owerId"`
	Icon      string           `gorm:"column:icon"           json:"icon"`
	Type      uint             `gorm:"column:type"           json:"type"`
	Desc      string           `gorm:"column:desc"           json:"desc"`
	CreatedAt utils.CustomTime `gorm:"column:created_at"     json:"createdAtt"`
	UpdatedAt utils.CustomTime `gorm:"column:updated_at"     json:"updatedAt"`
}

// 设置表名称  默认的表明会带s  链接表名会变成users
func (g *GroupBasic) TableName() string {
	return "group_basic"
}

func (g *GroupBasic) CreateGroupBasic(c *gin.Context, req *model.CreateGroupBasicRequest) (err error) {
	// 创建数据库模型对象
	groupBasicModel := &GroupBasic{
		Name:   req.Name,
		OwerId: req.OwerId,
		Icon:   req.Icon,
		Type:   req.Type,
		Desc:   req.Desc,
	}

	result := common.Db.Create(groupBasicModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *GroupBasic) GetGroupBasicsList(c *gin.Context, req *model.GetGroupBasicListRequest) (err error, users []GroupBasic) {
	// 构建查询条件，排除分页字段
	query := common.Db

	// 根据请求参数动态构建查询条件
	if req.Name != "" {
		query = query.Where("name = ?", req.Name)
	}
	if req.OwerId > 0 {
		query = query.Where("ower_id = ?", req.OwerId)
	}
	if req.Type > 0 {
		query = query.Where("type = ?", req.Type)
	}

	// 执行查询
	result := query.Limit(req.PageSize).Offset(utils.GetPageOffset(req.Page, req.PageSize)).Find(&users)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, users
}

func (g *GroupBasic) GetGroupBasicsInfo(c *gin.Context, req *model.GroupBasicIdRequest) (err error, groupBasicInfo GroupBasic) {
	//context, err := common.GetGroupBasicFromContext(c)
	//if err != nil {
	//	return err, GroupBasicBasic{}
	//}
	//id:=context.ID
	result := common.Db.Where("id=?", req.ID).First(&groupBasicInfo)
	if result.Error != nil {
		return result.Error, GroupBasic{}
	}
	return nil, groupBasicInfo
}

func (g *GroupBasic) UpdateGroupBasic(c *gin.Context, req *model.UpdateGroupBasicRequest) (err error) {
	// 将字符串时间转换为common.CustomTime
	updates := map[string]interface{}{
		"Name":   req.Name,
		"OwerId": req.OwerId,
		"Icon":   req.Icon,
		"Type":   req.Type,
		"Desc":   req.Desc,
	}

	result := common.Db.Model(&GroupBasic{}).Where("id = ?", req.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *GroupBasic) DeleteGroupBasic(c *gin.Context, req *model.GroupBasicIdRequest) (err error) {
	// 创建一个带有主键的实例，让 GORM 知道要删除哪个记录
	groupBasic := GroupBasic{}
	groupBasic.ID = req.ID

	result := common.Db.Delete(&groupBasic)
	if result.Error != nil {
		return result.Error
	}
	// 检查是否有记录被删除
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到ID为%d的记录", req.ID)
	}

	return nil
}
