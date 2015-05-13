package main

import "fmt"


func main(){
	str := "Go言語"
	bytes := str[0:8]
	str2 := string(bytes)
	for i, c:= range str2 {
		if c == 0xFFFD {
			str2 = str2[i:]
			break
		} else {
			fmt.Printf("%c", c)
		}
	}

	fmt.Printf("\n")
}