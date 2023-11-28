package main

import "fmt"

// if variable name starts with capital letter, then it is publicly available through the project.
const LoginToken string = "jjdjhjhejh"

func main()  {
	var username string = "manav"
	fmt.Println(username);
	fmt.Printf("variable is of type : %T \n", username);

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn);
	fmt.Printf("variable is of type : %T \n", isLoggedIn);

	var smallVal uint8 = 255 // Gives Error if >255
	fmt.Println(smallVal);
	fmt.Printf("variable is of type : %T \n", smallVal);

	var smallFloat float32 = 255.45555555555555555555 // Gives 255.45555 upto 5 decimals
	fmt.Println(smallFloat);
	fmt.Printf("variable is of type : %T \n", smallFloat);

	var largeFloat float64 = 255.45555555555555555555 // Gives more precise data than float32
	fmt.Println(largeFloat);
	fmt.Printf("variable is of type : %T \n", largeFloat);

	// default values and aliases
	var anotherIntVariable int;
	fmt.Println(anotherIntVariable);
	fmt.Printf("variable is of type : %T \n", anotherIntVariable);

	var anotherStrVariable string;
	fmt.Println(anotherStrVariable);
	fmt.Printf("variable is of type : %T \n", anotherStrVariable);

	// declaring variables implicitly

	var newStr = "website" // not defining any type in advance
	fmt.Println(newStr);
	fmt.Printf("variable is of type : %T \n", newStr);

	// newStr = 3 // Gives Error if another type value is assigned

	// no var style 
	noOfUsers := 5 // Can't be used outside of main function. Use var keyword in that case.
	fmt.Println(noOfUsers);
	fmt.Printf("variable is of type : %T \n", noOfUsers);


	// public variable
	fmt.Println(LoginToken);
	fmt.Printf("variable is of type : %T \n", LoginToken);
}