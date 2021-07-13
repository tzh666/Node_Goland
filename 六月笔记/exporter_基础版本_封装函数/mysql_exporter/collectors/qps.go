package collectors

import (
	"database/sql"
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type QpsController struct {
	db   *sql.DB
	desc *prometheus.Desc
}

func NewQpsController(db *sql.DB) *QpsController {
	return &QpsController{
		db: db,
		desc: prometheus.NewDesc(
			"mysql_global_status_querics", // 指标名字
			"mysql global status Querics", // 帮助信息
			nil,                           // 可变labels
			nil,                           // 固定labels
		),
	}
}

// 写入信息
func (c *QpsController) Describe(desc chan<- *prometheus.Desc) {
	desc <- c.desc
}

// 采集
func (c *QpsController) Collect(metrice chan<- prometheus.Metric) {
	var (
		name  string
		count float64
	)
	// 采集数据,有数据库的连接才能执行sql
	if err := c.db.QueryRow("show variables where Variable_name = ?", "wait_timeout").Scan(&name, &count); err != nil {
		log.Fatal(err)
	}
	/*
		数据采集:
			desc
			类型
			值
			可变labels
			固定labels
	*/
	// 写入要选的数据类型 Const 其他的对应其他的类型
	metrice <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count)
}
