package Model

import (
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	// 好像只有int型的默认值才有效，string类型的默认值无效
	Model
	ID			 int	`gorm:"AUTO_INCREMENT"` // 只有int类型才会自增
	Name         string	`gorm:"not null;unique_index" json:"name"`
	Age          int	`gorm:"default:18" json:"age"`
	Birthday     time.Time	`json:"birthday"`
	Email        string  `gorm:"type:varchar(100)" json:"email"`
	Role         string  `gorm:"default:'user';size:255" json:"role"` // set field size to 255, 设置的默认值无效，既不会存储到数据库中，也不会更新到结构体中
	MemberNumber *string `gorm:"unique" json:"member_number"` // set member number to unique and not null
	Num          int     `gorm:"AUTO_INCREMENT" json:"num"` // set num to auto incrementable
	Address      string  `gorm:"index:addr" json:"address"` // create index with name `addr` for address
	IgnoreMe     int     `gorm:"-" json:"ignore_me"` // ignore this field
}

func (u *User) TableName() string {
	// 表名的多样性
	//if u.Role == "admin" {
	//	return "admin_users"
	//} else {
	//	return "users"
	//}

	// 设置为true时会禁用表名的多样性，User表的名字会默认为user
	// Config.DB.SingularTable(true)
	return "users"
}

// 通过BeforeCreate可以为主键赋值——主要用于主键没有设置自增的情况（主键的类型为string）
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	UUID, _ := uuid.NewV4()
	scope.SetColumn("ID", UUID.String())
	return nil
}

// 在hooks中修改字段的值
//func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
//	if pw, err := bcrypt.GenerateFromPassword(user.Password, 0); err == nil {
//		scope.SetColumn("EncryptedPassword", pw)
//	}
//}
