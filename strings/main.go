package main

import (
	"fmt"
	"strings"
	"reflect"
)

func main() {
	// modify a string; change to upper or lower case
	a := "CamelCase"

	firstChar := a[0:1]
	fmt.Println("firstChar is", firstChar)

	fmt.Println("camelCase is", strings.ToLower(firstChar)+a[1:])

	// delete an element from a string array
	//var list string
	list := []string{"a","b","c","d","e"}
	fmt.Println(list)
	for i,e := range list{
		if e  == "c"{
			fmt.Println("index: ",i)
			list = append(list[:i], list[i+1:]...)
			break
		}
	}
	fmt.Println(list)

	var b *[]string
	b = &[]string{"aaa","bbb"}
	update(b)
	fmt.Println(b)
	c := *b
	fmt.Printf("kind b %s, kind c %s\n", reflect.TypeOf(b).Kind(),reflect.TypeOf(c).Kind())
	fmt.Printf("TypeOf b %s, TypeOf c %s\n", reflect.TypeOf(b).String(),reflect.TypeOf(c).String())
	c = nil
	update(&c)
	fmt.Println("c: ",c)
}

func update(a *[]string){
	c := "ccc"
	*a = append(*a, c)
}