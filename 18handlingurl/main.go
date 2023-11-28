package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=jhfjdhjdjhfdj"

func main() {
	fmt.Println("Welcome to url handling.")
	fmt.Println(myUrl)

	//Parsing urls

	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("Type of query params are %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for key, val := range qparams {
		fmt.Printf("Key : %v, Value: %v\n", key, val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
