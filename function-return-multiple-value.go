package main

import "fmt"

func getFullName() (string, string) {
	return "Haqi", "Ramadhani"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
