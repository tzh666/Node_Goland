package collectors

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/prometheus/client_golang/prometheus"
)

/*
	Mysql的最大连接数--------> show variables where Variable_name = "max_connections";
	Mysql响应的最大连接数----> show global status like 'max_used_connections';
*/

type SlowQuericsController struct {
	// 组合结构体
	db   *sql.DB
	desc *prometheus.Desc
}

// new方法注册
func NewSlowQuericsController(db *sql.DB) *SlowQuericsController {
	return &SlowQuericsController{
		db: db,
		// NewDesc(fqName, help string, variableLabels []string, constLabels Labels)
		desc: prometheus.NewDesc(
			"mysql_global_status_slow_querics", // 指标名字
			"mysql global status slow Querics", // 帮助信息
			nil,                                // 可变labels
			nil,                                // 固定labels
		),
	}
}

// 描述信息,这里调用new方法注册把描述信息写入里面
func (c *SlowQuericsController) Describe(desc chan<- *prometheus.Desc) {
	// 描述信息
	desc <- c.desc
}

// 定义指标,定义了一个只写的管道,类型是prometheus.Metric
func (c *SlowQuericsController) Collect(metrice chan<- prometheus.Metric) {
	var (
		name  string
		count float64
	)
	// 采集数据,有数据库的连接才能执行
	if err := c.db.QueryRow("show global status where variable_name = ?", "queries").Scan(&name, &count); err != nil {
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
	metrice <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count)
}
