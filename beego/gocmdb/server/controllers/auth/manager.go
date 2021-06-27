package auth

import (
	"github.com/astaxie/beego/context"

	"github.com/imsilence/gocmdb/server/models"
)

// 一堆接口,等会自己实现。因为有两种认证方式,所以方法定义成接口,后面去实现即可！
// manage 接口的封装
type AuthPlugin interface {
	Name() string
	Is(*context.Context) bool
	IsLogin(*LoginRequiredController) *models.User
	GoToLoginPage(*LoginRequiredController)
	Login(*AuthController) bool
	Logout(*AuthController)
}

type Manager struct {
	plugins map[string]AuthPlugin
}

func NewManager() *Manager {
	return &Manager{
		plugins: map[string]AuthPlugin{},
	}
}

// 注册
func (m *Manager) Register(p AuthPlugin) {
	m.plugins[p.Name()] = p
}

func (m *Manager) GetPlugin(c *context.Context) AuthPlugin {
	for _, plugin := range m.plugins {
		if plugin.Is(c) {
			return plugin
		}
	}
	return nil
}

func (m *Manager) IsLogin(c *LoginRequiredController) *models.User {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		return plugin.IsLogin(c)
	}
	return nil
}

func (m *Manager) GoToLoginPage(c *LoginRequiredController) {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		plugin.GoToLoginPage(c)
	}
}

func (m *Manager) Login(c *AuthController) bool {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		return plugin.Login(c)
	}
	return false
}

func (m *Manager) Logout(c *AuthController) {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		plugin.Logout(c)
	}
}

// 创建一个默认的Manager
var DefaultManger = NewManager()
