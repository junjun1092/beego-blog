package models

import (
        _ "github.com/go-sql-driver/mysql"

        "github.com/astaxie/beego"
        "github.com/astaxie/beego/orm"
        "fmt"
)


type User struct {
        Id       int
        UserName string
        Name     string
        Password string
        Email    string
}

func RegisterDB(){
        mysqluser := beego.AppConfig.String("mysqluser")
        mysqldb := beego.AppConfig.String("mysqldb")
        //注册驱动
        orm.RegisterDriver("mysql",orm.DRMySQL)
        //注册默认数据库root:@/test?charset=utf8，密码为空格式
        orm.RegisterDataBase("default","mysql",mysqluser + ":@/" + mysqldb + "?charset=utf8")
        fmt.Println("连接成功")
        //注册model
        //orm.RegisterModel(new (User))

        orm.RunSyncdb("default", false, true)

}

func init()  {
       //orm.RegisterModel(new (User))
        orm.RegisterModelWithPrefix("blog_", new (User))
}