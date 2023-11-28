package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to our pizza app : ");

	fmt.Print("Please rate our pizza on the scale of 1 to 10 : ")
	reader := bufio.NewReader(os.Stdin);

	input, _ := reader.ReadString('\n')

	// numRating, err := strconv.ParseFloat(input, 64); // This will give the input value as 5\n
	// numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64); // This will give the input value as 5
	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64); // This will give the input value as 5


	if err != nil {
		fmt.Println(err)
	} else if numRating > 10 || numRating < 1{
		fmt.Println("Please rate between 1 to 10")
	} else {
		fmt.Printf("Thnks for rating : %v \n", numRating)
		fmt.Printf("Type of variable is %T", numRating)

	}

}