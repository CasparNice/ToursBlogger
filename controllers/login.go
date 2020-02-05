// login.go
package controllers

import (
	"blog/models"
	_ "blog/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:WeiZiLin@tcp(127.0.0.1:3306)/satomi?charset=utf8")
}

func (this *LoginController) Post() {
	// 获取浏览器传递的值，并去除两边的空格
	Name := strings.TrimSpace(this.GetString("username"))
	Pwd := strings.TrimSpace(this.GetString("userpwd"))

	if Name == "" {
		beego.Info("请输入用户名")
		this.TplName = "login.html"
		return
	} else if Pwd == "" {
		beego.Info("请输入密码")
		this.TplName = "login.html"
		return
	}
	//查询用户名和密码
	o := orm.NewOrm()
	user := models.User{}
	user.Name = Name
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("用户名登录失败！！！")
		return
	}

	// 判断密码是否一致
	if user.Password != Pwd {
		beego.Info("密码登录失败！！！")
		return
	}

	//登录成功，记住登录状态
	this.Data["IsLogin"] = true

	//跳转页面
	this.Redirect("/", 302) //跳转到 login.html
	return
}
