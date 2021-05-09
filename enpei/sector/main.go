package main

import (
	"fmt"
	"math"
)
type Param struct {
	angle	float64
	r		float64
	lenArc	float64
	area	float64
}
func main() {

	var data = []Param{
		{angle: 45,r: 11,},
		{angle: 85,r: 25,},
		{angle: 300,r: 15,},
		{angle: 110,r: 32,},
		{angle: 150,r: 55,},
		{angle: 200,r: 4.6,},
	}
	for _,d := range data{
		cal(&d)
		fmt.Printf("Area: %f		Length of Arc: %f\n", d.area, d.lenArc)
	}
}

func cal (data *Param){
	data.lenArc = data.angle/360*2*math.Pi*data.r
	data.area = data.angle/360*math.Pi*data.r*data.r
}