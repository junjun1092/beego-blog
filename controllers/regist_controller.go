package controllers

import (
        "github.com/astaxie/beego"
        "github.com/astaxie/beego/orm"
)

type RegistController struct{
	beego.Controller
}

func (this *RegistController) ToRegist(){
	this.TplName = "regist/regist.html"
}

func (this *RegistController) Regist(){
	userName := this.GetString("userName")
	name := this.GetString("name")
        	password := this.GetString("password")
          email := this.GetString("email")
	o := orm.NewOrm()
        user := new(User)
}