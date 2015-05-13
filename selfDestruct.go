package main

import "fmt"
import "time"

func main(){
	abort := make(chan bool)
	count := make(chan int)

	go cancel(abort)
	go countDown(count)

	for {
		select {
		case i:= <- count:
			if 0 == i {
				selfDestruct()
				return
			}
			fmt.Printf("%d seconds remaining\n",i)
		case a := <-abort:
			if a {
				fmt.Printf("Self destruct aborted\n")
			} else {
				selfDestruct()
			}
			return
		}
	}
}

func cancel(reply chan bool) {
	fmt.Printf("This program will self destruct,do you wish to cansel?\n")
    var r int
	fmt.Scanf("%c", &r)
	switch r {
	case 'y': reply <- true
	case 'Y': reply <- true
	default: reply <- false
	}
}

func countDown(count chan int) {
	for i:=10; i >= 0; i-- {
		count <- i
		time.Sleep(1000000000)
	}
}

func selfDestruct(){
	fmt.Printf("selfDestruct()\n")
}