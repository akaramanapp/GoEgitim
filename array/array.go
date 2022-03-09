package main

import "fmt"

var sayiArray [5]int
var sayiArray2 [5]int = [5]int{1, 2, 3, 4, 5}

func main() {
	sayiArray3 := make([]int, 4)

	// array icine deger ekleme
	sayiArray3[1] = 5

	fmt.Println(sayiArray)  //output: [0 0 0 0 0]
	fmt.Println(sayiArray2) //output: [1 2 3 4 5]
	fmt.Println(sayiArray3) //output: [0 5 0 0]

	// len fonksiyonu arrayin uzunlugunu verir
	fmt.Println(len(sayiArray3)) //output: 4
}
