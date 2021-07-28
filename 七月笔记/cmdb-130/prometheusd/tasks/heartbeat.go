package tasks

import (
	"prometheusd/client"
	"prometheusd/config"
	"time"
)

type HeartbeatTask struct {
	option *config.Option
	client *client.Client
}

// new函数
func NewHeartbeatTask(option *config.Option, client *client.Client) *HeartbeatTask {
	return &HeartbeatTask{option, client}
}

// 定时心跳检测
func (t *HeartbeatTask) Run() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		event := map[string]interface{}{
			"uuid": t.option.UUID,
		}
		t.client.Heartbeat(event)
		<-ticker.C
	}
}
