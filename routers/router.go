package routers

import (
	"beego-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Login;post:DoLogin")
	beego.Router("/logout",&controllers.LoginController{},"get:Logout")
	beego.Router("/toRegist",&controllers.RegistController{}, "get:ToRegist;post:DoRegist")
       //	beego.Router("/doRegist",&controllers.RegistController{},)
}
