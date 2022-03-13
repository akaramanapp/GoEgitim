package main

import "fmt"

// iki parametre alan bir function tanimi
func add(x int, y int) int {
	return x + y
}

func main() {

	// functionun cagirilmasi
	fmt.Println(add(42, 13))
}
