package main

import "fmt"

// key value degerler tutmak icin kullanabiliriz.
var map1 map[int]string

func main() {
	map1 = make(map[int]string)

	map1[0] = "Ali"
	map1[1] = "Kerim"

	fmt.Println(map1)

	// map1 dizisinin uzunlugu
	fmt.Println(len(map1))
}
