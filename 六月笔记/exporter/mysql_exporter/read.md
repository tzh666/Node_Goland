1、定义指标步骤
a、定义一个结构体
b、NewQpsController 实现一个New方法
c、实现两个接口
    Dsecribe  <- 写入信息
        desc <- c.desc

    Collect   <- 采集数据
    /*
		数据采集:
			采集目标 (prometheus.NewDesc)
			类型
			值 (自己写函数返回一个float64类型的数据，也就是自己要获取的数据)
			可变labels
			固定labels
	*/
    metrice <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, c.status("wait_timeout"))
d、日志处理
	logrus     --> 记录日志
	
	lumberjack --> 日志数据控制

	mysql_exporter --> 配置文件 viper包:
		ini
		json
		yml
		toml
		自定义格式

	MySQL 认证(prometheus支持的):
		Authorization: Bearer Token(随机32字符串)
		Authorization: Basic base64(admin:123456)

2、注册指标
3、暴露API
4、启动http服务
5、部署,二进制打包,写dockerfile,编排到k8s上即可
6、prometheus采集数据