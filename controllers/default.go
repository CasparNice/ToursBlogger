package controllers

import (
	_ "blog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MainController struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:WeiZiLin@tcp(127.0.0.1:3306)/satomi?charset=utf8")
}

func (c *MainController) Get() {
	c.Data["IsLogin"] = false
	c.TplName = "index.html"
}
