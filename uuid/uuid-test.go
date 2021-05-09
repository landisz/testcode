package main

import (
	"fmt"

	util "testcode-uuid/util"
)

func main() {

	u := util.UuidGen()
	
	fmt.Printf("UUIDv4: %s\n", u)

}

