package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	fmt.Println("123454")
	//this.Data["Website"] = "beego.me"
	this.Data["PageTitle"] = "首页"
	this.TplName = "index.html"
}
