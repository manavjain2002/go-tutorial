package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "Welcome to User Input";
	fmt.Println(welcome);

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Kindly give rating to our pizza : ")

	input, _ := reader.ReadString('\n');
	fmt.Println("Thanks for rating , ", input)
	fmt.Printf("Type of variable is %T", input)
}