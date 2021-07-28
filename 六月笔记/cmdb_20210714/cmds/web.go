package cmds

import (
	"fmt"

	"github.com/spf13/cobra"

	// 建议顺序：标准库，系统库，第三方库，本项目库，不同分组使用空行分割开。
	"cmdb/config"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 定义根(root)命令
var webCommand = &cobra.Command{
	Use:   "web",         // 命令的名字
	Short: "web console", // 帮助信息
	Long:  "web console", // 详细帮助信息
	RunE: func(cmd *cobra.Command, args []string) error {

		// 配置beego日志文件路径
		beego.SetLogger("file", `{"filename" : "logs/cmdb.log"}`)
		// 日志级别,平时就用debug,生产用Informational
		beego.SetLevel(beego.LevelDebug)
		// 提示行数
		beego.SetLogFuncCall(true)
		// 干掉控制台的日志打印
		beego.BeeLogger.DelLogger("console")

		// 缓存cache配置
		config.Init("file", `{
		"CachePath"      : "tmp/cache", 
		"FileSuffix"     : ".cache",
	    "EmbedExpiry"    : "60", 
		"DirectoryLevel" : "3"
		}`)

		// beego Debug模式开启
		orm.Debug = verbose

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/cmdb?charset=utf8mb4&loc=PRC&parseTime=true",
			beego.AppConfig.DefaultString("mysql::User", "root"),
			beego.AppConfig.DefaultString("mysql::Password", "123456"),
			beego.AppConfig.DefaultString("mysql::Host", "root"),
			beego.AppConfig.DefaultInt("mysql::Port", 3306))

		// beego orm
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", dsn)

		// 判断是否连接成功
		if db, err := orm.GetDB("default"); err != nil {
			return err
		} else if err := db.Ping(); err != nil {
			return err
		}

		// 启动beego
		beego.Run()
		return nil
	}, // 程序入口
}

// 把命令加到根命令下
func init() {
	rootCommand.AddCommand(webCommand)
}
