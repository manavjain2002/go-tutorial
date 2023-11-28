package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to slices study.")

	var fruitList = []string{};
	fmt.Printf("Type of Fruitlist is : %T \n", fruitList)

	fruitList = append(fruitList, "Apple", "Kiwi", "Mango", "Orange")

	fruitList = append(fruitList[1:4])

	fmt.Println("Fruitlist is : ", fruitList)
	fmt.Println("Length of Fruitlist is : ", len(fruitList))

	var vegetableList = [4]string{"potato", "beans", "cucumber", "ladyfinger"};

	fmt.Println("vegetableList is : ", vegetableList)
	fmt.Println("Length of vegetableList is : ", len(vegetableList))

	highscores := make([]int, 4)
	highscores[0] = 10
	highscores[1] = 100
	highscores[2] = 26
	highscores[3] = 405
	// highscores[4] = 505 // gives error

	highscores = append(highscores, 444,555, 666)
	fmt.Println("vegetableList is : ", highscores)
	fmt.Println("Length of highscores is : ", len(vegetableList))

	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println("highscoresList is : ", highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	// how to remove value from slics based on index

	var courses = []string{"react", "vue", "android", "go", "node"}
	indexToRemove := 2
	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...)
	fmt.Println("coursesList is : ", courses)

}