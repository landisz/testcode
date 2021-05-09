package main

import (
	"fmt"
	"time"
)

func main() {
	var done = false
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			c1 <- "from one"
			time.Sleep(3 * time.Second)
		}
		c1 <- "quit"
	}()
	go func() {
		for i := 0; i < 15; i++ {
			c2 <- "from two"
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		select {
		case msg1 := <-c1:
			if msg1=="quit" {
				done = true
				break
			}
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
		if done {
			fmt.Println("Done")
			return
		}
	}
}

