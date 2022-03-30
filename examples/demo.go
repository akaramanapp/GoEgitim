package main

import "encoding/json"

func main() {
	user := Person{
		Name:    "abdulkerim",
		Surname: "karaman",
	}
	d := topla(1, 19)
	userArr, _ := json.Marshal(user)
	println(string(userArr))
	println(d)
}

func topla(deger1 int, deger2 int) int {
	return deger1 + deger2
}

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
