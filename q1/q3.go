package main

import "fmt"

// SOME type
type SOME struct {
	int
}

func main() {
	var x = SOME{}
	fmt.Printf("%#v\t%T\n", x, x)
	x.int = 42
	fmt.Printf("%#v\t%T\n", x, x)
}
