package main

import "fmt"

func main() {
	var name string

	name = "Haqi"
	fmt.Println(name)

	name = "Haqi Ramadhani"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age int8 = 22
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Haqi"
		lastName  = "Ramadhani"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
