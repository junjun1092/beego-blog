package controllers
import (
	"github.com/astaxie/beego"
        "github.com/astaxie/beego/orm"
        "beego-blog/models"
)
type LoginController struct{
	beego.Controller
}
func(this *LoginController) Login(){
	this.TplName = "login/login.html"
}
func (this *LoginController) DoLogin(){
	name := this.GetString("userName")
	if name == "" {
		this.Ctx.WriteString("userName is empty")
		return
	}
	password:= this.GetString("password")
	if password == ""{
		this.Ctx.WriteString("password is empty")
		return
	}
        models.RegisterDB()
        o := orm.NewOrm()
        o.Using("beego")
        user := new(models.User)
        user.UserName = name
        err := o.Read(user, "UserName")
        if err == orm.ErrNoRows {
	      this.Ctx.WriteString("没有查到相应的用户")
        }else {
	      pass := user.Password
	      if password == pass {
		    this.Ctx.WriteString("登陆成功")
	      }else {
		    this.Ctx.WriteString("密码错误")
	      }
        }
	this.Ctx.SetCookie("bb_name",name,2592000,"/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie","bb_password=" + password + "; Max-Age=2592000;Path=/;httponly")
	//this.Data["json"] =
	//this.ServeJSON()
        // this.Ctx.WriteString("登陆成功12")

}
func (this *LoginController) Logout(){
	this.Ctx.SetCookie("bb_name", "12", 0, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie","bb_password=" + "334" + "; Max-Age=0;Path=/;httponly")
	this.Ctx.WriteString("退出登陆112")
}

