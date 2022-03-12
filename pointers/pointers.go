package main

import "fmt"

// reference type
type String *string

func main() {
	var rstr String
	var pstr string

	// pointer'in zero value degeri nil'dir
	fmt.Println(rstr) // output: <nil>
	fmt.Println(pstr)

	pstr = "Hello"
	rstr = &pstr

	// degiskenin adresini gosteriyor
	fmt.Println(rstr) // output: 0xc000010230

	// eger degerine ulasmak istiyorsak
	fmt.Println(*rstr)

	// atadigimiz degeri gosteriyor
	fmt.Println(pstr) // output: Hello
}
