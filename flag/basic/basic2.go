package main

import "flag"
import "fmt"

func main() {
  var num = flag.Int("number", 1234, "test number flag")
	var str = flag.String("string", "hello", "test string flag")
	flag.Parse()
  fmt.Println("num value:", *num)
	fmt.Println("str value:", *str)
}
