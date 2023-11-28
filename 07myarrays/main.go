package main

import "fmt"

func main()  {
	fmt.Println("Welcome to arrays study.")

	var fruitList [4]string;

	fruitList[0] = "Apple"
	fruitList[1] = "Kiwi"
	fruitList[2] = "Mango"
	fruitList[3] = "Orange"

	fmt.Println("Fruitlist is : ", fruitList)
	fmt.Println("Length of Fruitlist is : ", len(fruitList))

	var vegetableList = [4]string{"potato", "beans", "cucumber", "ladyfinger"};

	fmt.Println("vegetableList is : ", vegetableList)
	fmt.Println("Length of vegetableList is : ", len(vegetableList))
}