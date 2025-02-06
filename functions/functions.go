package main

import (
	"fmt"
	"time"
)

func main_old() {
	x := 5
	y := func() int {
		return x * 2
	}
	fmt.Println("value of y is ", y())

	fmt.Println(" ::: ", time.Second)

	c := make(chan int)
	go func() {
		fmt.Println("Init Func")
		time.Sleep(5 * time.Second)
		fmt.Println("End Func")
		c <- 1
	}()
	<-c
}
