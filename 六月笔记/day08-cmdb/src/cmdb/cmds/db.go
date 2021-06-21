package cmds

// 建议顺序：标准库，系统库，第三方库，本项目库，不同分组使用空行分割开。
import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	db    string
	force bool
)

// 定义根(root)命令
var dbCommand = &cobra.Command{
	Use:   "db",         // 命令的名字
	Short: "db console", // 帮助信息
	Long:  "db console", // 详细帮助信息
	RunE: func(cmd *cobra.Command, args []string) error {
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
		// 从命令行获取参数,在init函数中绑定关系即可
		return orm.RunSyncdb(db, force, verbose)
	}, // 程序入口
}

// 把命令加到根命令下
func init() {
	rootCommand.AddCommand(dbCommand)
	/*
			&db,        要绑定的自定义参数地址
			"database", 长信息
			"d",        短信息 (只能是一个字符)
			"default",  默写
		    "database"  说明信息
	*/
	dbCommand.Flags().StringVarP(&db, "database", "d", "default", "database")
	dbCommand.Flags().BoolVarP(&force, "force", "f", false, "force")
}
