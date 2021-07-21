package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// how to take input that has spaces on that
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your full name : ")
	myname, _ := reader.ReadString('\n')
	fmt.Println(myname)

	// how to take an string input and then covert it to numeric value
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age in string : ")
	myrating, _ := reader2.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(mynumRating + 10)
}
