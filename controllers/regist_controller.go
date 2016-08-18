package controllers

import (
        "github.com/astaxie/beego"
       // "github.com/astaxie/beego/orm"
        "beego-blog/models"
        "fmt"
)

type RegistController struct{
	beego.Controller
}

func (this *RegistController) Get(){
	this.TplName = "regist/regist.html"
}

func (this *RegistController) DoRegist(){
        fmt.Println("123")
        userName := this.GetString("userName")
        name := this.GetString("name")
        password := this.GetString("password")
        email := this.GetString("email")
       // o := orm.NewOrm()
        user := new(models.User)
        user.UserName = userName
        user.Name = name
        user.Password = password
        user.Email = email
       // o.Insert(user)
        this.Ctx.WriteString("注册成功")
}