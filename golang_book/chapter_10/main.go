package main

import "fmt"
import "time"
import "math/rand"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)

		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func main(){
	fmt.Println("Let's start chapter 10")

	//go f(0)
	//
	//var input string
	//fmt.Scanf("%s", &input)

	//for i := 0; i < 10; i++ {
	//	go f(i)
	//}
	//
	//var input string
	//fmt.Scanln(&input)

	//var input string
	//
	//for i := 0; i < 10; i++ {
	//	go f(i)
	//	fmt.Scanln(&input)
	//}
	//
	//var c chan string = make(chan string)
	//
	//go pinger(c)
	//go ponger(c)
	//go printer(c)
	//
	//var input string
	//fmt.Scanln(&input)



}