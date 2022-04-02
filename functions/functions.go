package main

import "fmt"

// iki parametre alan bir function tanimi
func add(x int, y int) int {
	return x + y
}

// return degeri icermiyor
func addlog() {
	fmt.Println("Add log calisti.")
}

// tanimli degiskenleri return edebiliyoruz
func calculate(value int) (x, y int) { //return’un degiskenlerini tanımladık
	x = value / 2
	y = value * 2
	return // degiskenleri tanimladigimiz icin sadece return ediyoruz.
}

func main() {

	// panic example
	runServer()

	timer()

	// functionun cagirilmasi
	fmt.Println(add(42, 13))

	// return kullanmadik.
	addlog()

	// degisken tanimli fonksiyon
	fmt.Println(calculate(30))
}

func timer() {
	defer fmt.Println("Fonksiyonun sonunda bu calisacak")

	fmt.Println("Once bu kisim calisacak.")
}

func runServer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic olustu. Hata: ", err)
		}
	}()

	panic("")
	fmt.Println("Server calisti")
}
