package main

import (
"fmt"
"net"
"strconv"
)

func number() int {
	num := 15 * 5
	return num
}

func main() {

	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		//fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}

	_, cidr, _ := net.ParseCIDR("10.5.25.27/24")
	//subnet := string(cidr.IP)
	fmt.Println ("subnet is: ", cidr.IP)
	//fmt.Printf ("subnet is: %s", subnet )
	fmt.Printf ("subnet is: %v", cidr.IP[1] )
//	fmt.Printf ("subnet is: %T", cidr.IP[1] )
	s := fmt.Sprintf ("%v", cidr.IP)
	fmt.Println ("\ns is: ", s)
	var m uint8 = 24
	fmt.Println ("\ncidr  is: ", s+"/"+strconv.Itoa(int(m)))

}
