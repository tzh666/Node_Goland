package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

type CommandController struct {
	db   *sql.DB
	desc *prometheus.Desc
}

func NewCommandController(db *sql.DB) *CommandController {
	return &CommandController{
		db: db,
		desc: prometheus.NewDesc(
			"mysql_global_status_command",
			"mysql global status command",
			[]string{"command"}, // 可变标签
			nil,                 // 固定标签
		),
	}
}

func (c *CommandController) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *CommandController) Collect(metrics chan<- prometheus.Metric) {
	var (
		name  string
		count float64
	)
	/*
		数据采集:
			desc
			类型
			值
			可变labels
			固定labels
	*/
	// 写入要选的数据类型 Const 其他的对应其他的类型   <现在的labels就是可变的啦 insert-delete-update-select>
	c.db.QueryRow("show global status where Variable_name = ?", "Com_insert").Scan(&name, &count)
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count, "insert")

	// delete
	c.db.QueryRow("show variables where Variable_name = ?", "Com_insert").Scan(&name, &count)
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count, "delete")

	// update
	c.db.QueryRow("show variables where Variable_name = ?", "Com_insert").Scan(&name, &count)
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count, "update")

	// select
	c.db.QueryRow("show variables where Variable_name = ?", "Com_insert").Scan(&name, &count)
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count, "select")
}
