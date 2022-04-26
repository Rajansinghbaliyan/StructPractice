package main

func main() {
	james := person{
		"james",
		"bond",
		contactInfo{
			"james@gmail.com",
			201000}}

	harry := person{
		firstName: "harry",
		lastName:  "potter",
		contactInfo: contactInfo{
			email: "harrypotter@gmail.com",
			zip:   201010,
		},
	}

	james.print()
	harry.print()
}
