package test

import (
        "testing"
        "fmt"
        _"github.com/go-sql-driver/mysql"
        //"github.com/astaxie/beego/orm"
        "github.com/astaxie/beego/orm"
        "beego-blog/models"
)

//func init(){
//        fmt.Println("1")
//       //models.RegisterDB()
//}

func TestRegist(t *testing.T){
        orm.RegisterDriver("mysql",orm.DRMySQL)
        //注册默认数据库root:@/test?charset=utf8，密码为空格式
        orm.RegisterDataBase("default","mysql","root:@/beego?charset=utf8")
        orm.RegisterModelWithPrefix("prefix_", new (models.User))
        //models.RegisterDB()
        orm.RunSyncdb("default", false, true)
        o := orm.NewOrm()
        o.Using("beego")
        //user := models.User{Id:1}
        //var user models.User
        //user.Name = "Mary1"
        //user.UserName = "aa1"
        //user.Password = "1112"
        //user.Email = "1223@126.com"
        //fmt.Println(o.Insert(&user))
        //id, err := o.Insert(&user)
        //fmt.Println(err)
        //if err == nil {
        //        fmt.Println(id)
        //}
       // user := models.User{}
       //user.UserName = "aa"
       // err := o.Read(&user, "UserName")
       // if err == orm.ErrNoRows {
       //         fmt.Println("查询不到")
       // }else {
       //         fmt.Println("a:" ,user.Name)
       // }
}
