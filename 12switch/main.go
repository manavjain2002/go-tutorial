package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	fmt.Println("Welcome to switch study")

	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(6) + 1
	fmt.Println("Dicenumber : ", number)

	switch number {
	case 1:
		fmt.Println("Number is 1. Move 1 spaces")
	case 2:
		fmt.Println("Number is 2. Move 2 spaces")
	case 3:
		fmt.Println("Number is 3. Move 3 spaces")
	case 4:
		fmt.Println("Number is 4. Move 4 spaces")
		// fallthrough // if you want to execute the other cases as well mention the fallthrough in each case
	case 5:
		fmt.Println("Number is 5. Move 5 spaces")
	case 6:
		fmt.Println("Number is 6. Move 6 spaces and roll the dice again.")

	}

	// if err != nil {
	// 	fmt.Println(err)
	// } else if number == 0{
	// 	fmt.Println("Zero Number : ", number)
	// } else if number % 2 != 0{
	// 	fmt.Println("Odd Number : ", number)
	// } else {
	// 	fmt.Println("Even Number : ", number)
	// }

	// if num:=10; num <= 10 {
	// 	fmt.Println(num, "is Less than or equals to 10")
	// } else{
	// 	fmt.Println(num, "is Greater than 10")
	// }
}
