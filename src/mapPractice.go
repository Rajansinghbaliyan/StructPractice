package main

import "fmt"

func mapPractice() {
	countries := map[string]string{
		"india": "new delhi",
		"japan": "tokyo",
		"spain": "madrid",
	}
	fmt.Println(countries)

	for k, v := range countries {
		fmt.Printf("Key:%v, Value:%v", k, v)
		fmt.Println()
	}

	delete(countries, "india")

	fmt.Println(countries)

	capitalOfIndia := countries["india"]

	fmt.Println("Capital of india is: ", capitalOfIndia)

}
