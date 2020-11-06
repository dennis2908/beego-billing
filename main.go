package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	models "Learning-Beego/models"
	_ "github.com/go-sql-driver/mysql"
	_ "Learning-Beego/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/dennis?charset=utf8")
	orm.RegisterModel(new(models.Article),new(models.Company),new(models.C_bill))
	orm.RunCommand()
}

func main() {
	beego.Run()
}