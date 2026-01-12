package main

import "fmt"

func typeConversion() {
	a := 5
	fmt.Println(a)
	b := float32(a)
	fmt.Println(b)
	c := int32(b)
	fmt.Println(c)
	d := string(c)
	fmt.Println(d)
}
