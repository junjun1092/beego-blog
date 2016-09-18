package simpleTest

import(
	"fmt"
	"testing"
)

type  TZ int//底层类型，只能在相同包中绑定

func TestMethod(t *testing.T)  {
	//a := TZ()
	//a1:= 2
	var a TZ
	a.Print()
	//fmt.Print(a1)
	(*TZ).Print(&a)
}
func (a *TZ) Print(){
	fmt.Println("TZ")
}