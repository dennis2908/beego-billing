package main

import (
	models "Learning-Beego/models"
	_ "Learning-Beego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3307)/beego_billing?charset=utf8mb4")
	orm.RegisterModel(new(models.Article), new(models.Company), new(models.C_bill), new(models.User))
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
