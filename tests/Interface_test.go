package test

import (
	"fmt"
	"testing"
)

type  USB interface {
	Name() string
	Connector
}

type Connector interface {
	Connect()
}
type PhoneConnector struct {
	name string
} 

func(pc PhoneConnector) Name() string{
	return pc.name
}

func (pc PhoneConnector) Connect(){
	fmt.Println("Connected", pc.name)
}
func TestInt(t *testing.T){
	a := PhoneConnector{"PhoneConnector"}
	a.Connect()
	Disconnect(a)
	
}
func Disconnect(usb USB){
	//if pc,ok:= usb.(PhoneConnector); ok {
	//	fmt.Println("Disconnected",pc.name)
	//}
	switch  v:=usb.(type) {
	case
		PhoneConnector:
		fmt.Println("Disconnected", v.name)
	default:
		fmt.Println("Unkonwn device...")
	}


}