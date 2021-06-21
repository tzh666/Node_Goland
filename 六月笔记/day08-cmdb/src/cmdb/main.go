package main

import (
	_ "cmdb/routers"

	_ "github.com/go-sql-driver/mysql"

	"cmdb/cmds"
)

func main() {
	cmds.Execute()
}
