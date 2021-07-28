package services

import (
	"cmdb/models"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers() []*models.User {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&models.User{})
	var users []*models.User
	queryset.All(&users)
	return users
}

func AddUser(name string, password string, addr string, sex bool) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	user := &models.User{
		Name:     name,
		Password: string(hashed),
		Addr:     addr,
		Sex:      sex,
	}
	ormer := orm.NewOrm()
	ormer.Insert(user)
}

func DeleteUser(id int64) {
	ormer := orm.NewOrm()
	ormer.Delete(&models.User{ID: id})
}

func GetUserById(pk int64) *models.User {
	ormer := orm.NewOrm()
	user := &models.User{ID: pk}
	if err := ormer.Read(user); err == nil {
		return user
	}
	return nil
}

func GetUserByName(name string) *models.User {
	user := &models.User{Name: name}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

func Auth(username, password string) *models.User {
	// 通过用户名去查询用户信息(至少包含密码hash值)
	if user := GetUserByName(username); user == nil {
		// 用户名密码错误
		return nil
	} else {
		// 检查密码
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
			return user
		} else {
			// 用户名密码错误
			return nil
		}
	}
}
