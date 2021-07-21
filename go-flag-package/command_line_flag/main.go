package main

import (
	"flag"
	"fmt"
)

func main() {
	// https://gobyexample.com/command-line-flags
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("num", 42, "an integer")
	boolPtr := flag.Bool("bool", false, "a bool")

	var svr string
	flag.StringVar(&svr, "SVR", "bar", "a string var")

	flag.Parse()

	fmt.Println("word : ", *wordPtr)
	fmt.Println("num : ", *numPtr)
	fmt.Println("bool : ", *boolPtr)
	fmt.Println("SVR : ", svr)
	fmt.Println("tail : ", flag.Args())
}
