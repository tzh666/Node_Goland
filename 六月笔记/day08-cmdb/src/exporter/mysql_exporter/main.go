package main

import (
	"database/sql"
	"fmt"
	"mysql_exporter/collectors"
	"mysql_exporter/config"
	"mysql_exporter/handler"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	_ "github.com/go-sql-driver/mysql"
)

// 定义initLogger函数,返回值是func用于最后调用close(),相当于 defer logger.Close()
func initLogger(options config.Logger) func() {
	// 日志文件设置
	logger := lumberjack.Logger{
		Filename:   options.FileName,
		MaxSize:    options.MaxSize,    // 按大小,也可以按天数MaxDay
		MaxBackups: options.MaxBackups, // 保留天数
		Compress:   options.Compress,   // 日志压缩
	}
	// logrus的简单用法,日志格式配置
	logrus.SetOutput(&logger)                    // 输出到指定文件,这样所有logrus输出的日志都会写到文件中
	logrus.SetLevel(logrus.DebugLevel)           // 设置日志级别
	logrus.SetFormatter(&logrus.JSONFormatter{}) // 设置日志格式
	logrus.SetReportCaller(false)                // 控制是否打印日志代码行数,默认是false

	// 返回值是func用于最后调用close(),相当于 defer logger.Close()
	return func() {
		logger.Close()
	}
}

// 定义initDb函数,用于初始化数据库,返回值是*sql.DB
func initDb(options config.MySQL) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		options.UserName,
		options.PassWord,
		options.Host,
		options.Post,
		options.DbName,
	)
	return sql.Open("mysql", dsn)
}

// 定义initMetrics函数,用于初始化数据库,返回值是*sql.DB
func initMetrics(options *config.Options, db *sql.DB) {
	//  定义指标, 判断是否存活, mysql_up
	mysql_up := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "mysql_up",
			Help: "MySQL UP Info",
			ConstLabels: prometheus.Labels{"addr": fmt.Sprintf("%s:%d",
				options.MySQL.Host,
				options.MySQL.Post,
			)}, // 固定label
		},
		// 这里就是对应MySQL的数据了呗
		func() float64 {
			// ping的通就是存活,err == nil说明没有错误就是ping的通
			if err := db.Ping(); err == nil {
				return 1
			} else {
				// 打印日志且自定义日志描述
				logrus.WithFields(logrus.Fields{
					"metric": "mysql_info",
				}).Error(err)
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
	prometheus.MustRegister(collectors.NewConnertionController(db))
	prometheus.MustRegister(collectors.NewTrafficController(db))
}

func main() {

	options, err := config.ParseConfig("./etc/exporter.yaml")
	// 读取文件错误判断
	if err != nil {
		logrus.Error(err)
	}

	// 日志配置
	close := initLogger(options.Logger)
	defer close()

	// 连接数据库配置
	db, err := initDb(options.MySQL)
	if err != nil && db.Ping() != nil {
		logrus.Fatal(err)
	}

	// 初始化Metrics
	initMetrics(options, db)

	// 加密bcrypt.GenerateFromPassword
	// txt, _ := bcrypt.GenerateFromPassword([]byte("admin"), 5)
	// fmt.Println(string(txt))

	// 暴露API,暴露之前加一层闭包
	http.Handle(
		"/metrics/",
		handler.Auth(
			promhttp.Handler(),
			handler.AuthSecrets{
				options.Web.Auth.UserName: options.Web.Auth.PassWord,
			},
		))

	// 启动http服务
	http.ListenAndServe(options.Web.Addr, nil)
}
