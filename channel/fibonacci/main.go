package main

import (
	"fmt"
	"time"
)
func fibonacci(c, quit chan int) {
	fmt.Println("In fib func")
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			time.Sleep(1 * time.Second)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println("Got a num")
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	go fibonacci(c, quit)
}
