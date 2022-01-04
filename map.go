package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Haqi",
		"address": "Tegal",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Haqi Ramadhani"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
