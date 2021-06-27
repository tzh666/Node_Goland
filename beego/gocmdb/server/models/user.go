package models

import (
	"time"

	"github.com/imsilence/gocmdb/server/utils"

	"github.com/astaxie/beego/orm"
)

// 用户信息
type User struct {
	// column ---> 自定义数据库中的字段名称
	ID         int        `orm:"column(id)"`
	Name       string     `orm:"size(32)"`
	Password   string     `orm:"size(1024)"`
	Gender     int        `orm:"default(1)"`
	Brithday   *time.Time `orm:"null;default(null)"`
	Tel        string     `orm:"size(1024)"`
	Email      string     `orm:"size(1024)"`
	Addr       string     `orm:"size(1024)"`
	Remark     string     `orm:"size(1024)"`
	IsSuperman bool       `orm:"default(false)"`
	Status     int        `orm:"column(status)"`                   // 0正常 1锁定
	CreateTime *time.Time `orm:"column(create_time);auto_now_add"` // 一般时间都用指针
	UpdateTime *time.Time `orm:"column(update_time);auto_now"`
	DeleteTime *time.Time `orm:"column(delete_time);null;default(null)"`

	Token *Token `orm:"reverse(one)"`
}

// 对用户的操着都在的manager里面操着
type UserManager struct {
}

// 密码设置
func (u *User) SetPassword(password string) {
	u.Password = utils.Md5Salt(password, "")
}

// 密码校验,password是用户输入的,u.Password是通过id或者name去数据库中查出来的数据
func (u *User) ValidPassword(password string) bool {
	// 先获取到原来加密的盐
	salt, _ := utils.SplitMd5Salt(u.Password)
	// 再相同的盐加密用户输入的密码,加密去比较如果一致则认定是密码相同
	return utils.Md5Salt(password, salt) == u.Password
}

// 判断用户是否是锁定状态,从已经从数据库中查找出来的user信息去做对比
func (u *User) IsLock() bool {
	// 如果 u.Status == 1 说明是锁定的
	return u.Status == StatusLock
}

// 对用户的操着
// 根据id查找
func (m *UserManager) GetById(id int) *User {
	user := &User{}
	err := orm.NewOrm().QueryTable(user).Filter("ID__exact", id).Filter("DeleteTime__isnull", true).One(user)
	if err == nil {
		return user
	}
	return nil
}

// 根据name查找
func (m *UserManager) GetByName(name string) *User {
	user := &User{}
	err := orm.NewOrm().QueryTable(user).Filter("Name__exact", name).Filter("DeleteTime__isnull", true).One(user)
	if err == nil {
		return user
	}
	return nil
}

// 类似于一个构造函数,创建一个UserManager对象
func NewUserManager() *UserManager {
	return &UserManager{}
}

// token 认证
type Token struct {
	ID         int        `orm:"column(id)"`
	User       *User      `orm:"column(user);rel(one)"` // rel(oone)  一对一
	AccessKey  string     `orm:"column(access_key);size(1024)"`
	SecrectKey string     `orm:"column(secrect_key);size(1024)"`
	CreateTime *time.Time `orm:"column(create_time);auto_now_add"`
	UpdateTime *time.Time `orm:"column(update_time);auto_now"`
}

// token认证manager
type TokenManager struct {
}

// 数据库中验证是否有这个token
func (m *TokenManager) GetByKey(accessKey, secrectKey string) *Token {
	token := &Token{AccessKey: accessKey, SecrectKey: secrectKey}
	ormer := orm.NewOrm()
	// 到数据库中查找是否有这个key
	if err := ormer.Read(token, "accessKey", "secrectKey"); err == nil {
		// fmt.Println("token1", token.User)  // 现在只能查出来 id，其他的都是nil
		ormer.LoadRelated(token, "User") // 关联查询 token User
		//	fmt.Println("token2", token.User) // 这样就能查询出对token对应的User的信息
		return token
	}
	return nil
}

func NewTokenManager() *TokenManager {
	return &TokenManager{}
}

// 定义一个默认的Manager,赋值一个UserManager对象
var DefaultUserManager = NewUserManager()
var DefaultTokenManager = NewTokenManager()

// 注册模型
func init() {
	orm.RegisterModel(&User{})
	orm.RegisterModel(&Token{})
}
