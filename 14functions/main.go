package main

import "fmt"

func main() {
	fmt.Println("Welcome to function study.")
	greeter()

	fmt.Println("Adding 2 and 3 is : ", adder(2,3))

	values := []int{1,2,3,4,5,6,7,8,9,10}

	sum1, len1 := proAdder(values...)
	sum2, len2 := proAdder(1,2,3,4,5,6,7,8,9,10)

	fmt.Println("Addition of 1st 10 integers is : ", sum1, len1)
	fmt.Println("Addition of 1st 10 integers is : ", sum2, len2)
}

func greeter() {
	fmt.Println("Namaste")
}

func adder(value1 int, value2 int) int {
	return value1 + value2
}

func proAdder(values... int ) (int, int) {
	total := 0

	for _,value := range(values){
		total += value
	}

	return total, len(values)
}
