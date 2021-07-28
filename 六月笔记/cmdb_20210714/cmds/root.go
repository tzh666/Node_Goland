package cmds

import (
	"github.com/spf13/cobra"
	// 建议顺序：标准库，系统库，第三方库，本项目库，不同分组使用空行分割开。
	_ "cmdb/routers"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// 全局变量
var (
	verbose bool
)

// 定义根(root)命令
var rootCommand = &cobra.Command{
	Use:   "cmdb",          // 命令的名字
	Short: "cmdb programe", // 帮助信息
	Long:  "cmdb programe", // 详细帮助信息
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("cmdb")
		return nil
	}, // 程序入口
}

func init() {
	// 全局的verbose
	rootCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose")
}

// 执行方法
func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
