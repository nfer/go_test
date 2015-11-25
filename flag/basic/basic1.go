package main

import "flag"
import "fmt"

func main() {
	var num = flag.Int("number", 1234, "test number flag")
	flag.Parse()
	fmt.Println("num value:", *num)
}
