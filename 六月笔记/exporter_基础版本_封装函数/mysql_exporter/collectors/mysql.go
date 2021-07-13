package collectors

import (
	"database/sql"
)

type mysqlController struct {
	db *sql.DB
}

func (c *mysqlController) status(name string) float64 {
	sql := "show global status where Variable_name = ?"
	var (
		vname string
		rs    float64
	)
	if err := c.db.QueryRow(sql, name).Scan(&vname, &rs); err == nil {
		return rs
	}
	return 0
}

func (c *mysqlController) variables(name string) float64 {
	sql := "show global status where Variable_name = ?"
	var (
		vname string
		rs    float64
	)
	if err := c.db.QueryRow(sql, name).Scan(&vname, &rs); err == nil {
		return rs
	}
	return 0
}
