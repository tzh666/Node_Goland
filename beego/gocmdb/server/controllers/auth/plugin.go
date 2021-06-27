package auth

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"

	"github.com/imsilence/gocmdb/server/forms"
	"github.com/imsilence/gocmdb/server/models"
)

// Session结构体,底下方法实现AuthPlugin接口
type Session struct {
}

func (s *Session) Name() string {
	return "session"
}

// 判断验证方式,A、B两种，如果不是A那就是B
func (s *Session) Is(c *context.Context) bool {
	// 判断如果session等于空,就认定为token登录
	return c.Input.Header("Authentication") == ""
}

// 判断是否登录
func (s *Session) IsLogin(c *LoginRequiredController) *models.User {
	// session存的是一个id
	if session := c.GetSession("user"); session != nil {
		// 如果uid不等于nil，就用断言sid去数据库查这条数据
		if uid, ok := session.(int); ok {
			return models.DefaultUserManager.GetById(uid)
		}
	}
	return nil
}

// 跳转
func (s *Session) GoToLoginPage(c *LoginRequiredController) {
	c.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)
}

// 登录
func (s *Session) Login(c *AuthController) bool {
	form := &forms.LoginForm{}
	valid := &validation.Validation{}
	if c.Ctx.Input.IsPost() {
		// 解析页面form表单参数
		if err := c.ParseForm(form); err != nil {
			// 如果解析失败,打印下错误
			valid.SetError("error", err.Error())
		} else {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
			} else if ok {
				// 如果登录成功,那就把User.ID存在session里面
				c.SetSession("user", form.User.ID)
				c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
				return true
			}
		}
	}
	c.TplName = "auth/login.html"
	c.Data["form"] = form
	c.Data["valid"] = valid
	return false
}

// 退出登录
func (s *Session) Logout(c *AuthController) {
	c.DelSession("user")
	c.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)
}

// ******************************************Token**************************************************************
type Token struct {
}

func (t *Token) Name() string {
	return "token"
}

// 判断验证方式,A、B两种，如果不是A那就是B
func (t *Token) Is(c *context.Context) bool {
	return strings.ToLower(strings.TrimSpace(c.Input.Header("Authentication"))) == "token"
}

// 判断是否登录
func (t *Token) IsLogin(c *LoginRequiredController) *models.User {
	// 从请求头中获取accesskey和secrectkey
	accsesskey := strings.TrimSpace(c.Ctx.Input.Header("AccessKey"))
	secrectKey := strings.TrimSpace(c.Ctx.Input.Header("SecrectKey"))
	// toekn存在且不是被删除状态
	if token := models.DefaultTokenManager.GetByKey(accsesskey, secrectKey); token != nil && token.User.DeleteTime == nil {
		return token.User
	}
	return nil
}

// 跳转,token 不需要跳转.
func (t *Token) GoToLoginPage(c *LoginRequiredController) {
	c.Data["json"] = map[string]interface{}{
		"code":   403,
		"text":   "请使用正确的token进行认证",
		"result": nil,
	}
	c.ServeJSON()
}

// token不需要登录
func (t *Token) Login(c *AuthController) bool {
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "请使用Token请求API",
		"result": nil,
	}
	c.ServeJSON()
	return false
}

// 退出登录
func (s *Token) Logout(c *AuthController) {
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "退出登录成功",
		"result": nil,
	}
	c.ServeJSON()
}

func init() {
	DefaultManger.Register(new(Session))
	DefaultManger.Register(new(Token))
}
