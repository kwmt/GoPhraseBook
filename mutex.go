package main

import "fmt"
import "sync"

func main(){
	m := make(map[int] string)
	m[2] = "First Value"

	var lock sync.Mutex
	lock.Lock()
	go func(){
		m[2] = "Second Value"
//		lock.Unlock()
	}()
	
	lock.Lock()
	v:=m[2]
	lock.Unlock()
	fmt.Printf("%s\n", v)

}