package db

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
) // import your used driver

func init() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
}
