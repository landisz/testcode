package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := rand.Intn(20)
	i := rand.Intn(20)
	var b int
	fmt.Println("Calculate ", a, "-", i, "=")
	fmt.Scan(&b)
	if b == a - i {
		fmt.Println("correct. well done!")

	} else {
		fmt.Println("wrong, try again.")
	}

}