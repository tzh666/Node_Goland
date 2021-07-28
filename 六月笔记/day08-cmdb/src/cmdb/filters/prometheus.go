package filters

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/context"

	"github.com/prometheus/client_golang/prometheus"
)

// 定义以下标签
// url总请求次数      Counter    label随意
// 每个url请求次数    Counter    且带可变Label
// 状态码次数         Counter    且带可变Label
// 每个url的请求时间  Historgam   且带可变Label

var (
	// url总请求次数
	totalRequest = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cmdb_request_total",
		Help: "",
	})
	// 每个url请求次数
	urlRequest = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cmdb_url_request_total",
		Help: "",
	},
		// 可变label
		[]string{"url"},
	)
	// 状态码次数
	statusCode = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cmdb_status_code_total",
		Help: "",
	},
		[]string{"code"},
	)
	// 每个url的请求时间
	elapsedTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "cmdb_elapsed_time_request",
		Help: "",
	},
		[]string{"elapsedtime"},
	)
)

// 注册
func init() {
	prometheus.MustRegister(totalRequest, urlRequest, statusCode, elapsedTime)
}

// 过滤器机制实现要统计的数据类型
// 实现了四种数据类型统计更新,触发时间
func BeferExec(ctx *context.Context) {
	// 请求总数+1
	totalRequest.Inc()
	// 各url总请求+1,获取url(ctx.Input.URL()),可变label(WithLabelValues)
	urlRequest.WithLabelValues(ctx.Input.URL()).Inc()
	// url请求时间,起点
	ctx.Input.SetData("stime", time.Now())
}

// 过滤器机制实现要统计的数据类型
// 实现了四种数据类型统计更新,触发时间
func AfterExec(ctx *context.Context) {
	// 各状态码总次数,获取状态码(ctx.ResponseWriter.Status)
	statusCode.WithLabelValues(strconv.Itoa(ctx.ResponseWriter.Status)).Inc()
	stime := ctx.Input.GetData("stime")
	if stime != nil {
		if t, ok := stime.(time.Time); ok {
			// url请求时间,终点
			elapsed := time.Now().Sub(t)
			elapsedTime.WithLabelValues(ctx.Input.URL()).Observe(float64(elapsed))
		}
	}
}
