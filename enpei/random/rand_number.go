package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NumberOfQuestions = 1

func exchange(first int, second int) (int, int){
	return second, first
}

func showHello(){
	fmt.Println("Hello")
}
func main(){

	var a,b,c int
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	for i := 0; i < NumberOfQuestions; i++{
		a = rand.Intn(100)
		c = rand.Intn(100)
		if a < c {
			a,c = exchange(a, c)
		}

		fmt.Println("Question", i+1, ": ", a, "-", c, "=")
		fmt.Scan(&b)
		if b != (a - c) {
			fmt.Println("wrong, maybe next time.")
		} else {
			fmt.Println("correct. well done!")
		}
	}

	timeCost:= time.Since(start)
	k,_ := fmt.Println("You spent,",int(timeCost/1000000000),"seconds, to finish these questions, well done!!")
	fmt.Println(k)
	showHello()
}
