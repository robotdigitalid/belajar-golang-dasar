package main

import "fmt"

//goland:noinspection ALL
func main() {
	name := "Haqi"

	switch name {
	case "Haqi":
		fmt.Println("Hello Haqi")
	case "Joko":
		fmt.Println("Hello Joko")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Kenalan dong!")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length2 := len(name)
	switch {
	case length2 > 10:
		fmt.Println("Nama terlalu panjang")
	case length2 > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
