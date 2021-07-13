package main

import (
	"database/sql"
	"log"
	"mysql_exporter/collectors"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 连接数据库
	dsn := "root:123456@tcp(192.168.1.208:3306)/tzh666"
	mysqlAddr := "192.168.1.208:3306"
	addr := ":10001"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(db.Ping())
	//  定义指标, 判断是否存活, mysql_up
	mysql_up := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name:        "mysql_up",
			Help:        "MySQL UP Info",
			ConstLabels: prometheus.Labels{"addr": mysqlAddr}, // 固定label
		},
		// 这里就是对应MySQL的数据了呗
		func() float64 {
			// ping的通就是存活,err == nil说明没有错误就是ping的通
			if err := db.Ping(); err == nil {
				return 1
			}
			// 不通就返回0
			return 0
		},
	)

	// 注册指标
	prometheus.MustRegister(mysql_up)
	prometheus.MustRegister(collectors.NewSlowQuericsController(db))
	prometheus.MustRegister(collectors.NewQpsController(db))
	prometheus.MustRegister(collectors.NewCommandController(db))

	// 更新采样值
	// 暴露API
	http.Handle("/metrics/", promhttp.Handler())

	// 启动http服务
	http.ListenAndServe(addr, nil)
}
