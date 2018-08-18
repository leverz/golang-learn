package main

import (
	"fmt"
	"flag"
	"os"
	"../test1"
	test12 "../test2/test1"
	)

var name string

func init()  {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("hello, %s!\n", name)
	test1.Hello()
	test12.Hello()
}

