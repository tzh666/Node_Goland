package routers

import (
	"cmdb/controllers"
	"cmdb/controllers/api"
	"cmdb/filters"

	"github.com/astaxie/beego"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {

	// 过滤器注册到router
	beego.InsertFilter("/*", beego.BeforeExec, filters.BeferExec)
	beego.InsertFilter("/*", beego.AfterExec, filters.AfterExec, false)

	beego.ErrorController(&controllers.ErrorController{})

	// 暴露metrics接口,暴露后就可以直接访问了 (https://ip:10443/metrics)
	beego.Handler("/metrics/", promhttp.Handler())

	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.UserController{})

	beego.AutoRouter(&api.PrometheusController{})
}
