package main

import "flag"

func main() {
	var num = flag.Int("number", 1234, "test number flag")
	flag.Parse()
	println("num value:", *num)
}
