package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

// 流量监控
type TrafficController struct {
	mysqlController
	desc *prometheus.Desc
}

func NewTrafficController(db *sql.DB) *TrafficController {
	return &TrafficController{
		mysqlController: mysqlController{db},
		desc: prometheus.NewDesc(
			"mysql_globl_status_traffic", // name
			"mysql globl status traffic", // help
			[]string{"direction"},        // 可变标签
			nil,                          // 固定标签
		),
	}
}

func (c *TrafficController) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *TrafficController) Collect(metrics chan<- prometheus.Metric) {
	// desc结构体定义哪个,类型,返回值,标签拼接[direction="received"]
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, c.status("Bytes_received"), "received")
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, c.status("Bytes_sent"), "sent")
}
