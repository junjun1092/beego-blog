package test

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//第一个参数是数据库的别名，用来切换数据库
	//第二个是driverName,在RegisterDriver时注册的，第三个是数据库连接字符串
	//第四个参数相当于orm.SetMaxIdleConns("default", 30)。设置数据库的最大空闲连接
	//第五个参数相当于orm.SetMaxOpenConns("default", 30)，设置数据库的最大数据库连接，第四个和第五个可以不传，会使用数据库默认值
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30, 30)
	fmt.Println("连接成功")
}
