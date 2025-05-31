package main

import (
	"fmt"
	"time"
)

func testChannel() {
	c := make(chan int)
	go func() {
		fmt.Println("goroutine 1 execute")
		defer fmt.Println("goroutine 1 end")
		//time.Sleep(3 * time.Second)
		c <- 888
	}()
	time.Sleep(3 * time.Second)
	// channel作为同步点，会两边相互等待
	num := <-c
	fmt.Println("num = ", num)
	time.Sleep(2 * time.Second)
	fmt.Println("main goroutine end")
}
