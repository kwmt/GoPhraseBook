package main

import (
	"fmt"
	"time"
)

func main() {
	var a [100]int

	subrange := a[1:10]
//	for i,v:=range subrange {
//		fmt.Printf("Element: %d %d\n",i,v)
//	}

	for i,v:=range subrange{
		go fmt.Printf("Element: %d %d\n",i,v)
	}
	time.Sleep(10000000)

	
}