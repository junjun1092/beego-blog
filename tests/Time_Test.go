package test

import (
        "time"
        "fmt"
        "testing"
)

func TestTime(t *testing.T)  {
        t1 := time.NewTimer(2 * time.Second)
        go onTime(t1.C)
        fmt.Println("main thread")
        time.Sleep(10 * time.Second)
}

func onTime(c <- chan time.Time){
        for now := range c{
	      fmt.Println("OnTime", now)
        }
}


