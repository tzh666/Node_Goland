package config

import "github.com/spf13/viper"

// Web 信息
type Web struct {
	Addr string `mapstructure:"addr"`
	Auth struct {
		UserName string `mapstructure:"username"`
		PassWord string `mapstructure:"password"`
	} `mapstructure:"auth"`
}

// 解析yaml文件的值到对象中,调用的时候就可以直接调用对象的属性即可
type MySQL struct {
	Host     string `mapstructure:"host"`
	Post     int    `mapstructure:"port"`
	UserName string `mapstructure:"username"`
	PassWord string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
}

// 解析yaml文件的值到对象中,调用的时候就可以直接调用对象的属性即可
type Logger struct {
	FileName   string `mapstructure:"filename"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxSize    int    `mapstructure:"max_siza"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
}

// 定义结构体
// 解析yaml文件的值到对象中,调用的时候就可以直接调用对象的属性即可
type Options struct {
	MySQL  MySQL  `mapstructure:"mysql"`
	Web    Web    `mapstructure:"web"`
	Logger Logger `mapstructure:"logger"`
}

// 文件读取函数
func ParseConfig(path string) (*Options, error) {

	// new一个viper对象,viper 提供默认 Viper对象,可直接使用。也通过 New 方法创建自定义Viper
	conf := viper.New()

	// 默认值设置
	viper.SetDefault("mysql.port", 3306)
	viper.SetDefault("web.addr", ":10001")

	// 设置文件读取路径
	conf.SetConfigFile(path)

	// 判断读取文件是否有误
	if err := conf.ReadInConfig(); err != nil {

		// 判断配置文件是否存在,默认值都写里面
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			// 开启环境变量读取
			viper.AutomaticEnv()

			// 设置环境变量前缀,有这个前缀的环境变量才会读取 _连接
			viper.SetEnvPrefix("mysql_exporter")
		}
		return nil, err
	}

	// 初始化对象
	options := &Options{}

	// 判断解析文件到对象中是否有误
	// 解析yaml文件的值到对象中,调用的时候就可以直接调用对象的属性即可
	if err := conf.Unmarshal(options); err != nil {
		return nil, err
	}

	// 都ok则返回一个对象
	return options, nil
}
