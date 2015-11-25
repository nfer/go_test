# flag basic test

## flag basic test 1
Firstly, we want know what a basic flag usage looks like:

```
package main

import "flag"

func main() {
  var num = flag.Int("number", 1234, "test number flag")
  flag.Parse()
  println("num value:", *num)
}
```

after `go build basic1.go`, we get the executable `basic1`, below is the test and output:
```
$ ./basic1
num value: 1234
$ ./basic1 --help
Usage of ./basic1:
  -number int
      test number flag (default 1234)
$ ./basic1 -number 999
num value: 999
```

According to the test output, we can get below points:

1. without any flags(parameter), "number" keeps the defalut value
2. with `--help` flag, the executable output the help message
3. with `-number 999` or `-number=999`, "number" use the input value 999

## flag basic test 2
From last test, we know with `--help` flag, the executable output the help message, but is there any hint for help message?

```
$ ./basic1 -help
Usage of ./basic1:
  -number int
      test number flag (default 1234)
$ ./basic1 -h
Usage of ./basic1:
  -number int
      test number flag (default 1234)
$ ./basic1 --h
Usage of ./basic1:
  -number int
      test number flag (default 1234)
$ ./basic1 --help1
flag provided but not defined: -help1
Usage of ./basic1:
  -number int
      test number flag (default 1234)
$ ./basic1 --h1
flag provided but not defined: -h1
Usage of ./basic1:
  -number int
      test number flag (default 1234)
```

According to the test output, we know, only the `-help`, `--help`, `-h` and `--h` four flags are valid help message hint. 

## flag basic test 3
As `-help` flag works the same with `--help`, how it does on '-number' flags?
```
$ ./basic1 -number 999
num value: 999
$ ./basic1 --number 999
num value: 999
```

According to the test output, we knonw, go parsing flags -xx or --xx the same.
The go office doc page says:
> Flag parsing stops just before the first non-flag argument ("-" is a non-flag argument) or after the terminator "--".

## flag basic test 4
In last test, we only test one flag, how it works with multi flags, is it really "--" is a terminator?

The src code are below:
```
package main

import "flag"

func main() {
  var num = flag.Int("number", 1234, "test number flag")
  var str = flag.String("string", "hello", "test string flag")
  flag.Parse()
  println("num value:", *num)
  println("str value:", *str)
}
```

The test steps and outputs are below:
```
$ ./basic2 --number 999 --string "world"
num value: 999
str value: world
$ ./basic2 --number 999 -string "world"
num value: 999
str value: world
$ ./basic2 -number 999 -string "world"
num value: 999
str value: world
$ ./basic2 -number 999 --string "world"
num value: 999
str value: world
$ ./basic2 --number 999 -- --string "world"
num value: 999
str value: hello
$ ./basic2 --number 999 - --string "world"
num value: 999
str value: hello
```

According to the test output, we can get below points:

1. `-xx` and `--xx` work same on flag value, you can use `-xx` or `--xx` or mixed
2. `-xx` and `--xx` work same on terminator, the PURPOSE is to stop parsing last flags after the separate `-` or `--`.
