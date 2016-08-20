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
        orm.RegisterDriver("mysql",orm.DRMySQL)
        //注册默认数据库
        orm.RegisterDataBase("default","mysql","root:@/beego?charset=utf8")
        orm.RunSyncdb("default", false, true)
}

//func init()  {
//       orm.RegisterModel(new (User))
//        fmt.Println("regist user")
//}