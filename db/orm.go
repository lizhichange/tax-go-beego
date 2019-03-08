package db

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
) // import your used driver

func init() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)

	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	url := beego.AppConfig.String("mysqlurls")
	db := beego.AppConfig.String("mysqldb")

	_ = orm.RegisterDataBase("default", "mysql", ""+user+":"+pass+"@tcp("+url+":3306)/"+db+"?charset=utf8", 30)
}
