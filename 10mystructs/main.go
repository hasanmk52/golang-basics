package main

import "fmt"

func main() {
	fmt.Println("Structs in golang!")
	// no inheritance in golang; no super or parent

	hasan := User{"Hasan", "hasan@go.dev", true, 20}
	fmt.Println(hasan)
	fmt.Printf("Hasan details are: %+v\n", hasan)
	fmt.Printf("Name is %v and Email is %v.\n", hasan.Name, hasan.Email)
}

type User struct {	
	Name string
	Email string
	Status bool
	Age int
}