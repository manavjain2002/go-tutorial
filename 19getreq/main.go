package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	fmt.Println("Welcome to get request study")
	// PerformGetReq()
	// PerformPostJsonReq()
	PerformPostFormReq()
}

func PerformGetReq(){
	const myUrl = "http://localhost:3000"

	res, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Response status is : ", res.StatusCode)
	fmt.Println("Response length is : ", res.ContentLength)

	content, err := io.ReadAll(res.Body)
	fmt.Println(content)
	fmt.Println(string(content))

	var responseString strings.Builder;
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is : ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonReq()  {
	const myUrl = "http://localhost:3000/"

	requestBody := strings.NewReader(`
		{
			"courseName" : "golang",
			"price": 0,
			"platform": "xyz.com"
		}
	`)

	res, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body);
	fmt.Println(content)
	fmt.Println(string(content))
}

func PerformPostFormReq(){
	const muUrl = "http://localhost:3000/"

	data := url.Values{}
	data.Add("firstname", "Manav")
	data.Add("lastname", "Jain")

	res, err := http.PostForm(muUrl, data)
	if err!=nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}