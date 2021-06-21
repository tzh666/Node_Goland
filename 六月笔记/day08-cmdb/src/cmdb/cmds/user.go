package cmds

// 建议顺序：标准库，系统库，第三方库，本项目库，不同分组使用空行分割开。
import (
	"cmdb/models"
	"cmdb/utils"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	name     string
	password string
)

// 定义根(root)命令
var userCommand = &cobra.Command{
	Use:   "user",         // 命令的名字
	Short: "user console", // 帮助信息
	Long:  "user console", // 详细帮助信息
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

		// 初始化一个用户
		ormer := orm.NewOrm()
		user := &models.User{Name: name, Password: utils.GenerateFromPassword(password)}
		// 插入用户
		_, err := ormer.Insert(user)
		return err
	}, // 程序入口
}

// 把命令加到根命令下
func init() {
	rootCommand.AddCommand(userCommand)
	/*
			&db,        要绑定的自定义参数地址
			"database", 长信息
			"d",        短信息 (只能是一个字符)
			"default",  默写
		    "database"  说明信息
	*/
	userCommand.Flags().StringVarP(&name, "name", "n", "admin", "name")
	userCommand.Flags().StringVarP(&password, "password", "p", "admin", "password")
}
