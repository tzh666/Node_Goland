package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type ConnertionController struct {
	mysqlController
	maxConnertionDesc    *prometheus.Desc
	threadsConnectesDesc *prometheus.Desc
}

func NewConnertionController(db *sql.DB) *ConnertionController {
	return &ConnertionController{
		mysqlController: mysqlController{db},
		maxConnertionDesc: prometheus.NewDesc(
			"mysql_global_variabels_max_connectons",
			"mysql global variabels max connectons",
			nil,
			nil,
		),
		threadsConnectesDesc: prometheus.NewDesc(
			"mysql_global_status_threads_Connecte",
			"mysql global status threads Connecte",
			nil,
			nil,
		),
	}
}

func (c *ConnertionController) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.maxConnertionDesc
	descs <- c.threadsConnectesDesc
}

func (c *ConnertionController) Collect(metrics chan<- prometheus.Metric) {

	max_connections := c.variables("max_connections")
	// 打印日志且自定义日志描述
	logrus.WithFields(logrus.Fields{
		"metric": "max_connections",
		"sample": max_connections,
	}).Debug("command metric")

	// (desc *Desc, valueType ValueType, value float64, labelValues ...string)
	metrics <- prometheus.MustNewConstMetric(
		c.maxConnertionDesc,
		prometheus.CounterValue,
		max_connections,
	)

	metrics <- prometheus.MustNewConstMetric(
		c.threadsConnectesDesc,
		prometheus.CounterValue,
		c.status("threads_connected"),
	)
}
