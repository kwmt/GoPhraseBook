package main

import "fmt"

func main(){
	str := "Go言語"

	for i:=0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Printf("\n")

	for _,c := range str {
		fmt.Printf("%c",c)
	}
	fmt.Printf("\n")

}