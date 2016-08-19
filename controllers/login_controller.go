package controllers
import (
	"github.com/astaxie/beego"
	"fmt"
        "net/http"
)
type LoginController struct{
	beego.Controller
}
func(this *LoginController) Login(){
	this.TplName = "login/login.html"
}
func (this *LoginController) DoLogin(w http.ResponseWriter, r *http.Request){
	name := this.GetString("userName")
	fmt.Println("name:" + name)
	if name == "" {
		this.Ctx.WriteString("userName is empty")
		return
	}
	password:= this.GetString("password")
	if password == ""{
		this.Ctx.WriteString("password is empty")
		return
	}
	this.Ctx.SetCookie("bb_name",name,2592000,"/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie","bb_password=" + password + "; Max-Age=2592000;Path=/;httponly")
	this.Ctx.WriteString("登陆成功12")

}
func (this *LoginController) Logout(){
	this.Ctx.SetCookie("bb_name", "12", 0, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie","bb_password=" + "334" + "; Max-Age=0;Path=/;httponly")
	this.Ctx.WriteString("退出登陆112")
}

