package controllers

import (
        "github.com/astaxie/beego"
        "beego-blog/models"
        "github.com/astaxie/beego/orm"
        "fmt"
)

type RegistController struct{
	beego.Controller
}

func (this *RegistController) ToRegist(){
	this.TplName = "regist/regist.html"
}

func (this *RegistController) DoRegist(){
        userName := this.GetString("userName")
        name := this.GetString("name")
        password := this.GetString("password")
        email := this.GetString("email")
        models.RegisterDB()
        orm.Debug = true
        o := orm.NewOrm()
        o.Using("beego")
        user := new(models.User)
        user.UserName = userName
        user.Name = name
        user.Password = password
        user.Email = email
        id,err:= o.Insert(user)
        if err == nil {
                fmt.Println(id)
        }
        this.Data["json"] = map[string]interface{}{"success":0,"message":"注册成功"}
        this.ServeJSON()
        return
}