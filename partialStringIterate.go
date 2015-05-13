package main

import "fmt"
import "unicode/utf8"

func main(){
	str := "Go言語"
	rune := make([]byte, 0, 4)

	for i:=0; i < len(str); i++ {
		rune = append(rune, str[i])
		fmt.Printf("%x\n",rune)
		if utf8.FullRune(rune) {
			char, _ := utf8.DecodeRune(rune)
			fmt.Printf("%c", char)
			rune = rune[0:0]
			fmt.Printf("%v",i)
		}
	}
	fmt.Printf("\n")
}