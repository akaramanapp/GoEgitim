package main

import (
	"encoding/json"
	"fmt"
)

// User struct'i
type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {
	// User struct'ini kullanarak k objesi olusturuyoruz
	k := User{
		Name:    "Kerim",
		Surname: "Karaman",
	}

	// k struct'i json'a ceviriyoruz
	arr, _ := json.Marshal(k)
	fmt.Println(string(arr)) // {"name":"Kerim","surname":"Karaman"}
}
