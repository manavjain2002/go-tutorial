package main

import (
	"fmt"
	"sync"
)

func main()  {
	fmt.Println("WElcome to channel study")
	channel := make(chan int)
	wg := &sync.WaitGroup{}

wg.Add(2)
	go func(ch chan<- int, wg *sync.WaitGroup){
		// close(ch)
		
		ch <-10
		// ch<-11
		wg.Done()
	}(channel, wg)
	go func(ch <-chan int, wg *sync.WaitGroup){
		val, isChannelOpen := <-ch
		fmt.Println(val)
		fmt.Println(isChannelOpen)
		// fmt.Println(<-channel)
		// fmt.Println(<-channel)
		wg.Done()
	}(channel, wg)
	wg.Wait()
}