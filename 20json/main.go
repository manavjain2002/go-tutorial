package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"pricename"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to structs and json study")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	locCourses := []courses{
		{"ReactJs", 0, "1", "123@asc", []string{"web", "js", "android"}},
		{"VueJs", 1, "2", "12ddjnjd3@asc", []string{}},
		{"Android", 2, "3", "12ddjnjd3@asc", nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(locCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJs",
			"pricename": 0,
			"website": "1",
			"tags": [
					"web",
					"js",
					"android"
			]
		}
	`)

	var lcocourse courses

	checkValid := json.Valid(jsonDataFromWeb)
	fmt.Println(checkValid)

	if checkValid {
		fmt.Println("Valid Json")

		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you want to add data in key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, val := range(myOnlineData){
		fmt.Printf("Key : %v, Value: %v, Type : %T\n", key, val, val)
	}
}
