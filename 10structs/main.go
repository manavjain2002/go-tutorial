package main

import "fmt"

type User struct {
	Name string
	Email string
	Age int
	Status bool
}

func main()  {
	fmt.Println("Welcome to structs study")

	manav := User{"Mnav", "kddjekj", 16, false}
	fmt.Println(manav)
	fmt.Printf("Details are : %+v \n", manav)
	fmt.Printf("Name : %v, Email : %v", manav.Name, manav.Email)
}