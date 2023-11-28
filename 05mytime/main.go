package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome to our pizza app : ");
	presentTime := time.Now()

	fmt.Println("Recent order time is : ", presentTime)

	// To get date, fixed format is 01-02-2006
	// To get time, fixed format is 15:04:05
	// To get day, fixed format is Monday

	fmt.Println(presentTime.Local().Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(presentTime.UTC().Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(presentTime.Date())
	fmt.Println(presentTime.Nanosecond())
	fmt.Println(presentTime.Unix())

	createdDate := time.Date(2002, time.March, 27, 9, 9, 0, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Local().Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(createdDate.UTC().Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(createdDate.Date())
	fmt.Println(createdDate.Nanosecond())
	fmt.Println(createdDate.Unix())
}