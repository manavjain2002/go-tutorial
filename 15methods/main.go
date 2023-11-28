package main

import "fmt"

type User struct {
	Name string
	Email string
	Age int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("User status is : ", u.Status)
}

func (u User) SetEmail(mail string){
	oldMail := u.Email
	u.Email = mail

	fmt.Println(oldMail, "is changed to : ", u.Email)
}

func main()  {
	fmt.Println("Welcome to structs study")

	manav := User{"Mnav", "kddjekj", 16, true}
	fmt.Println(manav)
	fmt.Printf("Details are : %+v \n", manav)
	fmt.Printf("Name : %v, Email : %v\n", manav.Name, manav.Email)

	manav.GetStatus()
	manav.SetEmail("mndnjendj@jdjdjh")
	fmt.Printf("Name : %v, Email : %v\n", manav.Name, manav.Email)

}

