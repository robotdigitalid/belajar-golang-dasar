package main

import "fmt"

func getFullName2() (firstName string, middleName string, lastName string) {
	firstName = "Haqi"
	middleName = ""
	lastName = "Ramadhani"
	return
}

func main() {
	firstName, middleName, lastName := getFullName2()
	fmt.Println(firstName, middleName, lastName)
}
