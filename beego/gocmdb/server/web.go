package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/imsilence/gocmdb/server/models"
	"github.com/imsilence/gocmdb/server/utils"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/imsilence/gocmdb/server/routers"
)

func main() {

	// 初始化命令行参数
	h := flag.Bool("h", false, "help")
	help := flag.Bool("help", false, "help")
	init := flag.Bool("init", false, "init user")
	syncdb := flag.Bool("syncdb", false, "syncdb mysql")
	force := flag.Bool("force", false, "force")
	verbose := flag.Bool("verbose", false, "verbose")

	// 提示信息
	flag.Usage = func() {
		fmt.Sprintln("usage: web -h")
		flag.PrintDefaults()
	}
	// 解析命令行参数
	flag.Parse()
	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	// 日志设置
	beego.SetLogger("file",
		`{"filename" : "logs/web.log",
	      "level" : 7}`,
	)

	// 	// 取消控制台日志打印
	// 	beego.BeeLogger.DelLogger("console")

	// 初始化orm
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dsn"))

	// 测试数据量连接
	if db, err := orm.GetDB(); err != nil || db.Ping() != nil {
		beego.Error("数据库连接错误")
		os.Exit(-1)
	}

	// 根据参数选择执行流程
	switch {
	case *init:
		orm.RunSyncdb("default", *force, *verbose)
		// 插入一条数据
		ormer := orm.NewOrm()
		admin := &models.User{Name: "admin", IsSuperman: true}
		// 先判断数据库是否有这条数据
		// 意思是查询数据库Name字段是否有 Name=admin的字段,orm.ErrNoRows 空
		if err := ormer.Read(admin, "Name"); err == orm.ErrNoRows {
			password := utils.RandString(8)
			admin.SetPassword(password)
			// 插入数据
			if _, err := ormer.Insert(admin); err == nil {
				beego.Informational("初始化admin用户成功,默认密码是:", password)
			} else {
				beego.Error("初始化用户失败,错误: %s", err)
			}
		} else {
			beego.Informational("用户admin存在,跳过")
		}
	case *syncdb:
		// default 是注册数据库时候的别名 参数二为是否开启创建表  参数三是否更新表
		orm.RunSyncdb("default", *force, *verbose)
		// 日志记录
		beego.Informational("同步数据库成功")
	default:
		// 启动beego
		beego.Run()
	}
}
