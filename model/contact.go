package model

type CreateContactRequest struct {
	OwerId   uint   `gorm:"column:ower_id"        json:"owerId"   form:"owerId"   binding:"required"`
	TargetId uint   `gorm:"column:target_id"      json:"targetId" form:"targetId" binding:"required"`
	Type     uint   `gorm:"column:type"           json:"type"     form:"type"     binding:"required"`
	Desc     string `gorm:"column:desc"           json:"desc"     form:"desc"     binding:"omitempty"`
}

type GetContactListRequest struct {
	OwerId   uint `gorm:"column:ower_id"        json:"owerId"   form:"owerId"   binding:"omitempty"`
	TargetId uint `gorm:"column:target_id"      json:"targetId" form:"targetId" binding:"omitempty"`
	Type     uint `gorm:"column:type"           json:"type"     form:"type"     binding:"omitempty"`
	PageModel
}
type ContactIdRequest struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
}
type UpdateContactRequest struct {
	ID       uint   `gorm:"primaryKey;column:id"  json:"id"       form:"id"       binding:"required"`
	OwerId   uint   `gorm:"column:ower_id"        json:"owerId"   form:"owerId"   binding:"required"`
	TargetId uint   `gorm:"column:target_id"      json:"targetId" form:"targetId" binding:"required"`
	Type     uint   `gorm:"column:type"           json:"type"     form:"type"     binding:"required"`
	Desc     string `gorm:"column:desc"           json:"desc"     form:"desc"     binding:"omitempty"`
}
