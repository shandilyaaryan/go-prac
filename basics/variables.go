package main

import "fmt"

var python, java, golang bool

func variables() {
	var typescript bool = false
	python, java = false, false
	golang = true
	fmt.Printf("Python: %t, Java: %t, Typescript: %t, Go: %t\n", python, java, typescript, golang)
}
