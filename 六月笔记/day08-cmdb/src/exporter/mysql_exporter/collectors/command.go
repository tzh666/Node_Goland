package collectors

import (
	"database/sql"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type CommandController struct {
	mysqlController
	desc *prometheus.Desc
}

func NewCommandController(db *sql.DB) *CommandController {
	return &CommandController{
		mysqlController: mysqlController{db},
		desc: prometheus.NewDesc(
			"mysql_global_status_command",
			"mysql global status command",
			[]string{"command", "tzh"}, // 可变标签
			nil,                        // 固定标签
		),
	}
}

func (c *CommandController) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *CommandController) Collect(metrics chan<- prometheus.Metric) {
	/*
		数据采集:
			desc
			值类型  (那四种类型)
			值
			可变labels
			固定labels
	*/
	// 写入要选的数据类型 Const 其他的对应其他的类型   <现在的labels就是可变的啦 insert-delete-update-select>
	// insert
	// delete
	// update
	// select
	names := []string{
		"insert",
		"delete",
		"update",
		"select",
	}
	for _, name := range names {
		sample := c.status(fmt.Sprintf("Com_%s", name))
		// 打印日志且自定义日志描述
		logrus.WithFields(logrus.Fields{
			"metric": name,
			"sample": sample,
		}).Debug("command metric")
		metrics <- prometheus.MustNewConstMetric(
			c.desc,
			prometheus.CounterValue,
			sample,
			name,
			name,
		)
	}
}
