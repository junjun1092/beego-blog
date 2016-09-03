package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
        //	name := this.GetSession("name")
        //if name != nil {
	 //     this.Data["json"] = map[string]interface{}{"success":0,"message": name}
        //}
	this.TplName = "index.html"
}
