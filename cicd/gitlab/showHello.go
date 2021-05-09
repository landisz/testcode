package main

import (
	"fmt"
	"time"
)

const WHOM = "Ling"
func main() {
	for i:=0;i<10;i++{
		fmt.Println("Hello ", WHOM)
		time.Sleep(5 * time.Second)
	}
}

