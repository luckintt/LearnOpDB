package Model

import "github.com/guregu/null"

type Account struct {
	ID                 string      `gorm:"column:id;primary_key" json:"id"`
	Name               null.String `gorm:"column:name" json:"name"`
	Openid             null.String `gorm:"column:openid" json:"openid"`
	Unionid            string      `gorm:"column:unionid" json:"unionid"`
	Openid2            string      `gorm:"column:openid2" json:"openid2"`
	AreaCode           int         `gorm:"column:area_code" json:"area_code"`
	Mobile             string      `gorm:"column:mobile" json:"mobile"`
	Nickname           null.String `gorm:"column:nickname" json:"nickname"`
	Avatar             null.String `gorm:"column:avatar" json:"avatar"`
	AvatarCircle       null.String `gorm:"column:avatar_circle" json:"avatar_circle"`
	DeletedAt          null.Time   `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedAt          null.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          null.Time   `gorm:"column:updated_at" json:"updated_at"`
	Status             int         `gorm:"column:status" json:"status"`
	Kind               int         `gorm:"column:kind" json:"kind"`
	ClassTeacherID     null.Int    `gorm:"column:class_teacher_id" json:"class_teacher_id"`
	SalesID            null.Int    `gorm:"column:sales_id" json:"sales_id"`
	Idfa               null.String `gorm:"column:idfa" json:"idfa"`
	VerifyStatus       int         `gorm:"column:verify_status" json:"verify_status"`
	Type               int         `gorm:"column:type" json:"type"`
	XiaoshouyiID       string      `gorm:"column:xiaoshouyi_id" json:"xiaoshouyi_id"`
	MobileStatus       null.String `gorm:"column:mobile_status" json:"mobile_status"`
	PageSource         int         `gorm:"column:page_source" json:"page_source"`
	AcceptCreditCart   int         `gorm:"column:accept_credit_cart" json:"accept_credit_cart"`
	AccountType        null.Int    `gorm:"column:account_type" json:"account_type"`
	IsInternalEmployee int         `gorm:"column:is_internal_employee" json:"is_internal_employee"`
	CrmID              null.String `gorm:"column:crm_id" json:"crm_id"`
	IsAttendBargain    int         `gorm:"column:is_attend_bargain" json:"is_attend_bargain"`
	AccountSource      int         `gorm:"column:account_source" json:"account_source"`
	XiaoshouyiType     null.Int    `gorm:"column:xiaoshouyi_type" json:"xiaoshouyi_type"`
}

// TableName sets the insert table name for this struct type
func (a *Account) TableName() string {
	return "accounts"
}

