package main

import (
	"flag"
	"synchronous"
	"asynchronous"
)

var native bool

func main()  {
	flag.BoolVar(&native, "n", false, "using native method or not")
	flag.Parse()
	if native {
		synchronous.Run()
	} else {
		asynchronous.Run()
	}
}
