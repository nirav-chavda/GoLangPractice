package main

import "fmt"

func main() {
	p1 := struct {
		name string
		age  int
	}{
		name: "Nirav",
		age:  13,
	}

	fmt.Println(p1)

	// Anonymous Function
	func(x int) {
		fmt.Println(x)
	}(12)
}
