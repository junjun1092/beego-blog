package models

import ("github.com/astaxie/beego/orm"
        _ "github.com/go-sql-driver/mysql"
)


type User struct {
        Id       int
        UserName string
        Name     string
        Password string
        Email    string
}

func RegisterDB(){
        //注册model
        orm.RegisterModel(new (User))
        //注册驱动
        //orm.RegisterDriver("mysql",orm.or)
        //注册默认数据库
        orm.RegisterDataBase("default","mysql","root:@localhost/beego?charset=utf8")
}