package api

import (
	"cmdb/base/controllers"
	"cmdb/forms"

	"encoding/json"
	"fmt"
)

type PrometheusController struct {
	// 组合api的controller
	controllers.ApiController
}

func (c *PrometheusController) Alert() {
	// 一次copy数据大小
	c.Ctx.Input.CopyBody(1024 * 1024)
	// 接收Alertmanager发生过来的数据
	// c.Ctx.Input.RequestBody ==> 数据映射到结构体中(form)
	// 定义表,储存到数据库
	// fmt.Println(string(c.Ctx.Input.RequestBody))
	form := forms.AlterForm{}
	// Unmarshal 解析json数据
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err == nil {
		fmt.Println(err)
		for _, alter := range form.Alerts {
			fmt.Printf("%#v\n", alter)
		}
	}
	c.Data["json"] = map[string]string{"code": "200"}
	c.ServeJSON()
}
