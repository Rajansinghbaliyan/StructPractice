package main

import "fmt"

func main() {
	//unbuffered()
	buffered()
}

func unbuffered() {
	channel := make(chan int)

	// Will create a deadlock because there will be no listner
	//channel <- 2
	// These children will be blocked if there is no listener
	go func() { channel <- 1 }()
	go func() { channel <- 2 }()
	go func() { channel <- 3 }()

	first := <-channel
	second := <-channel
	third := <-channel

	fmt.Println(first, second, third)
}

func buffered() {
	channel := make(chan int, 2)

	channel <- 1
	channel <- 2

	// will also create a deadlock because the buffered capacity is 2
	//channel <- 3

	go func() { channel <- 4 }()

	first := <-channel
	second := <-channel
	third := <-channel

	fmt.Println(first, second, third)

}
