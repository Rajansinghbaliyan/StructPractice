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

	// This is the way to save the pointer
	harryPointer := &harry

	harryPointer.updateName("harry2")
	harryPointer.print()

	harry.updateName("harry3")
	harry.print()

	updateLastName(&harry, "potter2")

	harry.print()
}
