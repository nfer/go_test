package main

import "flag"

func main() {
  var num = flag.Int("number", 1234, "test number flag")
	var str = flag.String("string", "hello", "test string flag")
	flag.Parse()
  println("num value:", *num)
	println("str value:", *str)
}
