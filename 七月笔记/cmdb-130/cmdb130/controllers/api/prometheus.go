package api

import (
	"cmdb/base/controllers"
	"cmdb/forms"
	"cmdb/services"
	"encoding/json"
	"fmt"
)

type PrometheusController struct {
	controllers.ApiController
}

func (c *PrometheusController) Alert() {
	c.Ctx.Input.CopyBody(1024 * 1024 * 10)
	// c.Ctx.Input.RequestBody => 结构体
	// 定义表
	// 存储
	form := forms.AlertsForm{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err == nil {
		for _, alert := range form.Alerts {
			services.AlertService.Notify(alert.ToModel())
			fmt.Println(alert, "接收告警")
		}
	} else {
		fmt.Println(err)
	}

	c.Data["json"] = map[string]string{"code": "200"}
	c.ServeJSON()
}
