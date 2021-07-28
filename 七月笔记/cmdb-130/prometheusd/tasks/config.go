package tasks

import (
	"encoding/json"
	"prometheusd/client"
	"prometheusd/config"
	"prometheusd/domain"

	"time"
)

type ConfigTask struct {
	option *config.Option
	client *client.Client
}

// new函数
func NewConfigTask(option *config.Option, client *client.Client) *ConfigTask {
	return &ConfigTask{option, client}
}

// 定时配置文件更新
func (t *ConfigTask) Run() {
	ticker := time.NewTicker(time.Minute)
	var version int64
	for {
		event := map[string]interface{}{
			"uuid":    t.option.UUID,
			"version": version,
		}
		if txt, err := t.client.Config(event); err != nil {
			var result struct {
				Code    string
				Config  domain.Jobs
				Version int64
			}
			err := json.Unmarshal([]byte(txt), &result)
			if err == nil {
				version = result.Version
				// 更改配置文件
				if len(result.Config) > 0 {

				}
			}
		}
		<-ticker.C
	}
}
