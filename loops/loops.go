package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// dizi icindeki degerleri ekrana yazdiriyoruz
	for i, value := range arr {
		fmt.Println(i, value)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
