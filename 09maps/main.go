package main

import "fmt"

func main()  {
	fmt.Println("Welcome to maps study")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Languages list is : ", languages)
	fmt.Println("JS shorts for : ", languages["JS"])

	delete(languages, "RB")

	for key, value := range languages {
		fmt.Printf("Key : %v, Value: %v \n", key, value)
	}
}