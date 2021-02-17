package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	printArray(arr...)
	println([]interface{}{1, 2, 3, 4, 5})
}

func printArray(args ...int) {
	for _, i := range args {
		fmt.Println(i, ` Lol 

"end"`)
	}
}

func println(args ...interface{}) {
	for i := range args {
		fmt.Println(args[i], " end")
	}
}
