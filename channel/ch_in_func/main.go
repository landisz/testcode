package main

import (
	"fmt"
	"time"
)

type Tick struct {
	t *time.Ticker
	ch chan int
}

func main() {
	q := make(chan int)
	go prtHello(q)
	<-q
	fmt.Println("Done!")
}

func prtHello(quit chan int) {
	a := arrange(quit)
	i:=0
	//stopper(a.ch, quit)
	for {
		<-a.t.C
		Hello()
		if i>5 {a.ch<-0}
		i+=1
	}
}

func arrange(quit chan int) Tick{
	var a Tick
	a.t = time.NewTicker(2 * time.Second)
	a.ch = make(chan int)
	go stopper(a.ch, quit)
	return a
}

func Hello() {
	fmt.Println("Hello")
}

func stopper(c,q chan int){
	<-c
	fmt.Println("Stop now")
	q<-0
}