package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"
	str1 := "Istanbul"

	// ilk 5 karakteri ekrana bas
	fmt.Println(str[:5]) // output: Hello
	fmt.Println(str)     // output: Hello, World!

	fmt.Println(fmt.Sprintf("%s %s", str, str1)) // output: Hello, World! Istanbul

	// str degiskeni hello ile basliyor mu?
	if strings.HasPrefix(str, "Hello") {
		fmt.Println("Hello ile basliyor.")
	}

	fmt.Println(strings.ToUpper(str)) // output: HELLO, WORLD!
}
