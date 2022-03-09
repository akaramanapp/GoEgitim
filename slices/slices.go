package main

import "fmt"

var slice1 []int

// golang' da array"ler limitlidir, boyutlari bellidir. Slice"larin ise boyutlari genisleyebilir.

func main() {
	slice2 := make([]int, 0, 5)

	/*
	 dizinin boyutu sifir olarak tanimlandigi icin deger atamak istedigimizda hata aliriz.
	 slice2[0] = 6
	*/

	// append fonksiyonu ile slice'e eleman ekleriz.
	slice2 = append(slice2, 6) //output: [6]

	fmt.Println(slice1) //output: []
	fmt.Println(slice2) //output: []

	// cap fonksiyonu slice'in kapasitesini verir.
	fmt.Println(cap(slice2)) //output: 5
}
