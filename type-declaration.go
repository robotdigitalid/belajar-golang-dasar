package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpHaqi NoKTP = "093940283940423"
	var marriedStatus Married = false
	fmt.Println(noKtpHaqi)
	fmt.Println(marriedStatus)
}
