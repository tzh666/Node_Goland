package tasks

import (
	"prometheusd/client"
	"prometheusd/config"
	"time"
)

type RegistertTask struct {
	option *config.Option
	client *client.Client
}

// new函数
func NewRegistertTask(option *config.Option, client *client.Client) *RegistertTask {
	return &RegistertTask{option, client}
}

// 定时注册
func (t *RegistertTask) Run() {
	ticker := time.NewTicker(time.Hour)
	for {
		event := map[string]interface{}{
			"uuid":     t.option.UUID,
			"addr":     t.option.Addr,
			"hostname": t.option.Hostanem,
		}
		// 注册
		t.client.Register(event)
		// fmt.Println(event)
		<-ticker.C
	}
}
