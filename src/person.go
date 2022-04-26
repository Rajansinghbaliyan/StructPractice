package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func (p person) print() {
	//fmt.Printf("FirstName: %v LastName: %v", p.firstName, p.lastName)
	fmt.Printf("%+v", p)
	fmt.Println()
}
