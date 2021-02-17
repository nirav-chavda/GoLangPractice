package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%d %s %t\n", x, y, z)
	fmt.Printf("%d\n%s\n%t", x, y, z)
}
