package main

import "fmt"
import "time"

func main(){
	time.AfterFunc(200000000, func() {
		fmt.Printf("Timer expired\n")
	})
	timer := time.NewTimer(3000000000)
	fmt.Printf("Current time: %d nanoseconds\n", <- timer.C)

}