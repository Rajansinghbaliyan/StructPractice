package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

func (c contactInfo) print() {
	fmt.Printf("%+v", c)
	println()
}
