package controllers
import (
	"github.com/astaxie/beego"
        "github.com/astaxie/beego/orm"
        "beego-blog/models"
        "fmt"
)
type LoginController struct{
	beego.Controller
}
func(this *LoginController) Login(){
	this.TplName = "login/login.html"
}
func (this *LoginController) DoLogin(){
        name := this.GetString("userName")
        fmt.Println("name:" + name)
	if name == "" {
		fmt.Println( "userName is empty")
		return
	}
	password:= this.GetString("password")
	if password == ""{
		fmt.Println("password is empty")
		return
	}
        models.RegisterDB()
        o := orm.NewOrm()
        o.Using("beego")
        user := new(models.User)
        user.UserName = name
        err := o.Read(user, "UserName")
        fmt.Println(err)
        if err == orm.ErrNoRows {
	      this.Data["json"] = map[string]interface{}{"success":1,"message":"没有查到相应的用户"}
	      this.ServeJSON()
	      return
        }else {
	      pass := user.Password
	      if password == pass {
		    this.Data["json"] = map[string]interface{}{"success":0,"message":"登陆成功"}
		    this.ServeJSON()
		    beego.sess
		    return
	      }else {
		    fmt.Println("werrr")
		    this.Data["json"] = map[string]interface{}{"success":1,"message":"密码错误"}
		    this.ServeJSON()
		    return
	      }
        }
	this.Ctx.SetCookie("bb_name",name,2592000,"/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie","bb_password=" + password + "; Max-Age=2592000;Path=/;httponly")

}
func (this *LoginController) Logout(){
	this.Ctx.SetCookie("bb_name", "12", 0, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie","bb_password=" + "334" + "; Max-Age=0;Path=/;httponly")
	this.Ctx.WriteString("退出登陆112")
}

