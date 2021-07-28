package models

import (
	"cmdb/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

// 用户信息
type User struct {
	ID         int        `orm:"column(id)"`
	StaffID    string     `orm:"column(staff_id);size(32)"`
	Name       string     `orm:"size(64)"`
	Nickname   string     `orm:"size(64)"`
	Password   string     `orm:"size(1024)"`
	Gender     int        `orm:""`
	Tel        string     `orm:"size(128)"`
	Addr       string     `orm:"size(64)"`
	Email      string     `orm:"size(64)"`
	Department string     `orm:"size(128)"`
	Status     int        `orm:""`
	CreateAt   *time.Time `orm:"auto_now_add"` // 一般时间都用指针
	UpdateAt   *time.Time `orm:"auto_now"`
	DeleteAt   *time.Time `orm:"null"`
}

// 验证用户输入的密码是否正确
func (u *User) VaildPassword(password string) bool {
	// u.password数据库的密码跟获取用户输入的密码做比较
	return utils.CheckPassword(u.Password, password)
}

// 自定义函数,用户状态显示
func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "正常"
	case 1:
		return "锁定"
	case 2:
		return "离职"
	}
	return "用户状态错误"
}

// 注册模型
func init() {
	// 或者 orm.RegisterModel(new(User))
	orm.RegisterModel(&User{})
}
