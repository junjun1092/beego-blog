package main


import (
	//_ "beego-blog/routers"
	"fmt"
)

type user struct {
	Name string
	Age int
}

func main() {
	//beego.Run()
	//匿名结构
	//a := &struct {
	//	Name string
	//	Age int
	//}{
	//	Name : "Joe",
	//	Age:19,
	//}
	a := &user{Name:"Lucy",Age:19}
	A(a)
	fmt.Print(a)//Lucy 19，这里修改的是它的副本，没有修改原对象
}
func A(user *user){
	user.Name = "Lily"
}



