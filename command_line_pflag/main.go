package main

import flag "github.com/spf13/pflag"
import "fmt"


func main() {
	// https://github.com/spf13/pflag

	var s string
	fmt.Scanln(&s)

	wordPtr := flag.String("name", s, "is a name")
	numPtr := flag.Int("integer", 0, "an integer")
	boolPtr := flag.Bool("human", true, "a boolean")

	var svar string
	flag.StringVar(&svar, "word", "hello", "a word")

	flag.Parse()

	fmt.Println("name : ", *wordPtr)
	fmt.Println("number : ", *numPtr)
	fmt.Println("boolean : ", *boolPtr)
	fmt.Println("word : ", svar)
	fmt.Println("tail : ", flag.Args())

	fmt.Println("End of Parsing")
}