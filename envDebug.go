package main

import "fmt"
import "os"
import "strconv"

var debugLevel int

func debugLog(level int, msg string, args ...interface{}){
	if debugLevel > level { fmt.Printf(msg,args...) }
}

func main() {
	debugLevel, _ = strconv.Atoi(os.Getenv("DEBUG"))
	fmt.Printf("debugLevel=%d\n",debugLevel)
	debugLog(3, "Starting\n")
}