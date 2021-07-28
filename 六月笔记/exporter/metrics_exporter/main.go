package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 自定义结构体获取指标,请求的时候触发
type CpuController struct {
	cpuDesc *prometheus.Desc
}

func NewCpuController() *CpuController {
	return &CpuController{
		// NewDesc(fqName, help string, variableLabels []string, constLabels Labels)
		cpuDesc: prometheus.NewDesc(
			"test_cpu_percent_v2",
			"Cpu Percent V2",
			[]string{"cpu"}, // 可变label
			nil,             // 固定label
		),
	}
}

func (c *CpuController) Describe(descs chan<- *prometheus.Desc) {
	// 写数据
	descs <- c.cpuDesc
}

func (c *CpuController) Collect(metrics chan<- prometheus.Metric) {
	// 采集值
	// MustNewConstMetric(desc *Desc, valueType ValueType, value float64, labelValues ...string)
	for i := 0; i < 4; i++ {
		metrics <- prometheus.MustNewConstMetric(
			c.cpuDesc,             // 值
			prometheus.GaugeValue, // 类型
			rand.ExpFloat64(),     // 值类型
			strconv.Itoa(i),       // cpu name
		)
	}
}

func main() {
	// 定义指标,固定Labels
	totalV1 := prometheus.NewCounter(
		prometheus.CounterOpts{
			// Namespace,Subsystem,Name,这三个只写Name就行。因为他会用下划线自动帮我们拼接
			Namespace:   "",
			Subsystem:   "",
			Name:        "test_total_v1",         // Name  test_total_v1{name="v1"} 10
			Help:        "Test Total v1 Counter", // 帮助信息,会在页面显示
			ConstLabels: map[string]string{"name": "v1"},
		},
	)

	// 定义指标, 可变的Labels
	totalV2 := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "test",
			Subsystem:   "total",
			Name:        "v2",
			Help:        "Test Total v2 Counter",
			ConstLabels: prometheus.Labels{"name": "v2"},
		},
		[]string{"path"},
	)

	// 定义指标Counter,-->NewCounterFunc。这种方式-->请求发生时更新数据
	totalV3 := prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "test_total_v3",
		},
		func() float64 {
			fmt.Println("totalV3")
			return rand.Float64() // 这里就是指标采样值-->NewCounterFunc。这种方式-->请求发生时更新数据
		},
	)

	// 定义指标Gauge test_cpu_percent{cpu="/0/<就是这个0是可变的>"} 9.405090880450125！用的比较多
	cpuPercent := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "test_cpu_percent",
		Help: "Test CPU Percent",
	},
		[]string{"cpu"},
	)

	// 定义指标Histog,定义区间(每次间隔3)
	requestTimeH := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_time_H",
			Help:    "Request Time Histogrom",
			Buckets: prometheus.LinearBuckets(0, 3, 3), // 这里就是定义区间(每次间隔3) 0-3 0-6 0-9 0-Inf
		},
		[]string{"path"},
	)

	// 定义指标 Summary,request_time_S{path="/login/",quantile="0.7"} 17.939839151237454
	requestTimeS := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "request_time_S",
			Help:       "Request Time Summary",
			Objectives: map[float64]float64{0.9: 0.01, 0.8: 0.02, 0.7: 0.03}, // 0.9是百分比 0.01是偏差值
		},
		[]string{"path"},
	)

	// 注册指标
	prometheus.MustRegister(totalV1)
	prometheus.MustRegister(totalV2)
	prometheus.MustRegister(totalV3)
	prometheus.MustRegister(cpuPercent)
	prometheus.MustRegister(requestTimeH)
	prometheus.MustRegister(requestTimeS)

	// 自定义结构体指标注册
	prometheus.MustRegister(NewCpuController())

	// 更新指标采样值
	go func() {
		for range time.Tick(10 * time.Second) { // 延迟10秒
			totalV1.Add(10)                           // 定时更新,刚开始的时候是0
			totalV2.WithLabelValues("/root/").Inc()   // 可变的Labels
			totalV2.WithLabelValues("/login/").Add(5) // 可变的Labels

			cpuPercent.WithLabelValues("/0/").Set(rand.Float64() * 10)
			cpuPercent.WithLabelValues("/1/").Set(rand.Float64() * 10)

			requestTimeH.WithLabelValues("/root/").Observe(rand.Float64() * 20)
			requestTimeH.WithLabelValues("/login/").Observe(rand.Float64() * 20)

			requestTimeS.WithLabelValues("/root/").Observe(rand.Float64() * 20)
			requestTimeS.WithLabelValues("/login/").Observe(rand.Float64() * 20)
		}
	}()

	// 暴露 http api
	http.Handle("/metrics/", promhttp.Handler())

	addr := ":9009"
	// 启动web服务
	http.ListenAndServe(addr, nil)
}
