package test

import (
        "testing"
        "fmt"
        _"github.com/go-sql-driver/mysql"
        "github.com/astaxie/beego/orm"
        "beego-blog/models"
)

//func init(){
//       models.RegisterDB()
//
//}

func testRegist(t *testing.T){
        fmt.Println("regist")
       o := orm.NewOrm()
        o.Using("default")
        //user := models.User{Id:1}
        var user models.User
        user.Name = "Mary"
        user.UserName = "aa"
        user.Password = "111"
        user.Email = "123@126.com"
        id, err := o.Insert(user)
        if err == nil {
                fmt.Println(id)
        }
        //err := o.Read(&user)
        //if err == orm.ErrNoRows {
        //        fmt.Println("查询不到")
        //}else if err == orm.ErrMissPK{
        //        fmt.Println("找不到主键")
        //}else{
        //        fmt.Println(user.Id, user.Name)
        //}
        //o.Insert(user)
        //db, err := sql.Open("mysql","root:@/beego?charset=utf8")
        //checkErr(err)
        //stmt, err := db.Prepare("insert into user values(?,?,?,?,?)")
        //checkErr(err)
        //res, err := stmt.Exec(1,"Mary","mary","123","123@126.com")
        //checkErr(err)
        //id, err:= res.LastInsertId()
        //checkErr(err)
        //fmt.Println(id)
        //mysqluser := beego.AppConfig.String("mysqluser")
        //mysqldb := beego.AppConfig.String("mysqldb")
        //orm.RegisterDriver("mysql",orm.DRMySQL)
        //orm.RegisterDataBase("default","mysql",mysqluser + ":@/" + mysqldb + "?charset=utf8")
        //fmt.Println("链接成功")
}

//func checkErr(err error){
//        if err != nil {
//                panic(err)
//        }
//}