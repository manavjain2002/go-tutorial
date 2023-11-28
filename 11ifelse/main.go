package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	fmt.Println("Welcome to if else study")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number to check if it's even or odd : ")
	input, _ := reader.ReadString('\n')
	number, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if err != nil {
		fmt.Println(err)
	} else if number == 0{
		fmt.Println("Zero Number : ", number)
	} else if number % 2 != 0{
		fmt.Println("Odd Number : ", number)
	} else {
		fmt.Println("Even Number : ", number)
	}

	if num:=10; num <= 10 {
		fmt.Println(num, "is Less than or equals to 10")
	} else{
		fmt.Println(num, "is Greater than 10")
	}
}
