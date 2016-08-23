package controllers

import (
        "github.com/astaxie/beego"
       // "github.com/astaxie/beego/orm"
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
        user.Id = 1
        user.UserName = userName
        user.Name = name
        user.Password = password
        user.Email = email
        id,err:= o.Insert(user)
        if err != nil {
                fmt.Println(id)
        }
        this.Ctx.WriteString("注册成功12")
}