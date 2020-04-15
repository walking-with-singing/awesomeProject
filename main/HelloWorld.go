package main

import (
	"fmt"
	"time"
)

func main() {
	now:=time.Now()
	count:=0
	for i:=0;i<10000000000;i++{
		count+=i
	}
	fmt.Println("master")
	t2:=time.Now()
	tsub:=now.Sub(t2)
	fmt.Println(tsub)
	fmt.Println(int32(tsub.Seconds()))
}