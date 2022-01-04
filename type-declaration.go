package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpHaqi NoKTP = "3328181301990002"
	var marriedStatus Married = false
	fmt.Println(noKtpHaqi)
	fmt.Println(marriedStatus)
}
