package services

import (
	"cmdb/forms"
	"cmdb/models"
	"cmdb/utils"
	"fmt"
	"log"

	"github.com/astaxie/beego/orm"
)

type userService struct {
}

// 增加用户
func (s *userService) Add(form *forms.UserModifyForm) {

}

// 修改用户密码
func (s *userService) ModifyPassword(pk int, password string) {
	// 通过ID查询到User的信息,如果user!=nil,就通过orm修改密码
	if user := s.GetByPk(pk); user != nil {
		fmt.Println(user)
		user.Password = utils.GenerateFromPassword(password)
		orm.NewOrm().Update(user, "Password")
	}
}

// 通过ID删除用户
func (s *userService) Delete(pk int) {
	// 先创建一个User对象,然后初始化ID
	user := &models.User{ID: pk}
	// 获取数据库连接(获取数据库操作实例)
	orm.NewOrm().Delete(user)
}

// 编辑用户
func (s *userService) Modify(form *forms.UserModifyForm) {
	// 同样的先通过pk拿到一条用户的数据
	if user := s.GetByPk(form.ID); user != nil {
		// 把页面用户输入的参数赋值给原来的
		user.Name = form.Name
		user.Tel = form.Tel
		user.Addr = form.Addr
		user.Email = form.Email
		// 更新数据
		orm.NewOrm().Update(user, "Name", "Tel", "Addr", "Email")
	}
}

// 通过用户id获取一条用户信息
func (s *userService) GetByPk(pk int) *models.User {
	// 先创建一个User对象,然后初始化ID
	user := &models.User{ID: pk}
	// 获取数据库连接(获取数据库操作实例)
	ormer := orm.NewOrm()
	// 参数一：要读取的结构体的对象，参数二：按照ID去Read
	if err := ormer.Read(user, "ID"); err == nil {
		return user
	}
	return nil
}

// 通过用户名获取用户
func (s *userService) GetByName(name string) *models.User {

	user := &models.User{Name: name}
	// 获取数据库连接(获取数据库操作实例)
	ormer := orm.NewOrm()
	// 参数一：要读取的结构体的对象，参数二：按照Name去Read
	if err := ormer.Read(user, "Name"); err != nil {
		log.Fatal(err)
	} else {
		// 读取成功,直接返回用户信息
		fmt.Println(user)
		return user
	}
	// 否在返回nil
	return nil
}

// 查询用户,数据库返回的数据是切片(这是查询的逻辑) q是传入的查询条件
func (s *userService) Query(q string) []*models.User {
	// 定义一个切片，用户存放查询出来的数据
	var users []*models.User

	// 获取数据库连接(获取数据库操作实例),拿到数据库的查询数据集
	// ormer := orm.NewOrm()  queryset := ormer.QueryTable(&User{})  这2句可以直接写成下面的2句
	queryset := orm.NewOrm().QueryTable(&models.User{})
	// 有条件的时候的查询语句写法
	if q != "" {
		cond := orm.NewCondition()
		cond = cond.Or("name__icontains", q)
		cond = cond.Or("nickname__icontains", q)
		cond = cond.Or("tel__icontains", q)
		cond = cond.Or("addr__icontains", q)
		cond = cond.Or("email__icontains", q)
		cond = cond.Or("department__icontains", q)
		queryset = queryset.SetCond(cond) // 条件拼接 SetCond
	}
	// 没有查询添加的时候查询语句写法
	queryset.All(&users)
	return users
}

// UserService 创建用户操作服务(controller直接调用这里的Service)
var UserService = new(userService)
