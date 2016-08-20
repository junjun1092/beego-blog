package test

import (
        "testing"
        "beego-blog/models"
        "fmt"
        //"github.com/astaxie/beego/orm"
        "github.com/astaxie/beego/orm"
)

func init(){
       models.RegisterDB()

}

func testRegist(t *testing.T){
       o := orm.NewOrm()
        o.Using("default")
        user := models.User{Id:1}
        fmt.Println("123")
        //err := o.Read(&user)
        //if err == orm.ErrNoRows {
        //        fmt.Println("查询不到")
        //}else if err == orm.ErrMissPK{
        //        fmt.Println("找不到主键")
        //}else{
        //        fmt.Println(user.Id, user.Name)
        //}
        o.Insert(user)
}