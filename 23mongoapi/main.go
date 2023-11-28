package main

import (
	"fmt"
	"log"
	"mongoapi/routers"
	"net/http"
)

func main()  {
	fmt.Println("Welcome to mongo api learning.")
	r := routers.Router()
	fmt.Println("Server starting listening.......")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server listening on port 4000......")

}