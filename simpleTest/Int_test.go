package simpleTest

import (
	"fmt"
	"testing"
)
//没有方法的接口即为顶级父接口，所有结构都实现了这个接口
type empty interface {

}
type USB interface {
	Name() string
	//嵌入接口
	Connector
}

type  Connector interface {
	Connector()
}
type PhoneConnector struct {
	name string
}

func (phone PhoneConnector) Name() string{
	return phone.name
}
func (phone PhoneConnector) Connector(){
	fmt.Println("phone 已连接", phone.name)
}

func TestIntface(t *testing.T){
	//var usb USB
	var a Connector
	pc := PhoneConnector{"PhoneConnector"}
	a = Connector(pc)
	a.Connector()
	pc.name = "aaa"
	fmt.Println(pc.name)
	//usb.Connector()
	//fmt.Println(usb.Name())
	//DisConnect(usb)
}

func DisConnect(usb interface{}){
	//类型断言
	//if pc,ok := usb.(PhoneConnector);ok {
	//	fmt.Println("DisConnect", pc.name)
	//	return
	//}
	//fmt.Println("Unkown device")
	switch v:=usb.(type) {
	case PhoneConnector:
		fmt.Println("DisConnect", v.name)
	default:
		fmt.Println("Unkown device")
	}
}