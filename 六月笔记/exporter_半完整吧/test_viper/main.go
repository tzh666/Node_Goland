package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 配置文件解析到对象里面去
type Options struct {
	// 对应yaml的mysql  (Mysql--转成小写对应)
	MySQL struct {
		Host     string
		Password string
	}
	// 对应yaml的web
	Web struct {
		Auth map[string]string
	}
	// log
	Log struct {
		Max_age  int
		Max_size int
	}
	// 如果有不符合的就用标签去指定
	Test struct {
		Host string
	} `mapstructure:"db"`
}

func main() {
	// 指定要读取的配置文件
	viper.SetConfigType("yaml") // 指定文件格式
	viper.SetConfigName("test") // 文件名不用加后缀
	viper.AddConfigPath(".")    // 路径

	// 设置环境变量读取
	viper.AutomaticEnv()
	// 设置环境变量前缀,有这个前缀的环境变量才会读取 _连接
	viper.SetEnvPrefix("mysql_exporter")

	// 判断读取配置文件是否有误
	if err := viper.ReadInConfig(); err != nil {
		logrus.Debug(err)
	}

	// 设置默认值,配置文件不存在的时候才会用这个默认值。文件不存在也会用这个默认值
	viper.SetDefault("mysql.port", 3306)

	// 读取配置文件的值
	fmt.Println(viper.Get("mysql"))
	fmt.Println(viper.GetString("mysql.host"))
	fmt.Println(viper.GetInt("mysql.port"))

	// fmt.Println(viper.Get("redis.port")) // set redis.port=6333 就能读取到了 Linux就配置环境变量

	//
	options := new(Options)
	if err := viper.Unmarshal(options); err != nil {
		logrus.Error(err)
	}
	fmt.Println(options) // &{{192.168.1.208 123456}}

	// viper 写配置文件
	viper.SetDefault("redis.host", "2.2.2.2")
	viper.WriteConfigAs("./test02.yaml")
}
