package main

import (
	"fmt"
	"strings"
)

// Books s
type Books struct {
	title   string
	author  string
	subject string
	bookID  int
}

func main() {
	fmt.Println(Display())
	fmt.Println("Hello There")

	//var z,y = 10, "20"
	b := 10
	fmt.Println(b)

	// Expression Switch
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You passed\n")
	case grade == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}
	fmt.Printf("Your grade is  %s\n", grade)

	// Type switch
	var x interface{} // only with interface type

	switch i := x.(type) { // (type) is only used in this switch case
	case nil:
		fmt.Printf("type of x :%T", i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case func(int) float64:
		fmt.Printf("x is func(int)")
	case bool, string:
		fmt.Printf("x is bool or string")
	default:
		fmt.Printf("don't know the type")
	}

	var a int
	numbers := [6]int{1, 2, 3, 5}

	/* for loop execution */
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}
	for a < 10 {
		a++
		fmt.Printf("value of a: %d\n", a)
	}
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}

	//var p, q string
	p, q := swap("Mahesh", "Kumar")
	fmt.Println(p, q)

	// Strings
	fmt.Println(strings.Join([]string{"sample", "good"}, " "))

	arr := [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* output each array element's value */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("arr[%d][%d] = %d\n", i, j, arr[i][j])
		}
	}

	// Pointers

	var val = 20 /* actual variable declaration */
	var ip *int  /* pointer variable declaration */

	ip = &val /* store address of a in pointer variable*/

	fmt.Printf("Address of val variable: %x\n", &val)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	b = 12
	// Pass pointer as parameter
	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	swap_p(&a, &b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)

	// Structures
	var Book1 Books /* Declare Book1 of type Book */
	var Book2 Books /* Declare Book2 of type Book */

	/* book 1 specification */
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.bookID = 6495407

	/* book 2 specification */
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.bookID = 6495700

	/* print Book1 info */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.bookID)

	/* print Book2 info */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.bookID)

	// Slice
	var numbersSlice = make([]int, 3, 5)
	printSlice(numbersSlice)

	// Sub slicing
	numbersSlice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbersSlice)

	/* print the original slice */
	fmt.Println("numbers ==", numbersSlice)

	/* print the sub slice starting from index 1(included) to index 4(excluded)*/
	fmt.Println("numbers[1:4] ==", numbersSlice[1:4])

	/* missing lower bound implies 0*/
	fmt.Println("numbers[:3] ==", numbersSlice[:3])

	/* missing upper bound implies len(s)*/
	fmt.Println("numbers[4:] ==", numbersSlice[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* print the sub slice starting from index 0(included) to index 2(excluded) */
	number2 := numbersSlice[:2]
	printSlice(number2)

	/* print the sub slice starting from index 2(included) to index 5(excluded) */
	number3 := numbersSlice[2:5]
	printSlice(number3)

	numbersSlice = nil

	numbersSlice = append(numbersSlice, 0)
	printSlice(numbersSlice)

	/* add one element to slice*/
	numbersSlice = append(numbersSlice, 1)
	printSlice(numbersSlice)

	/* add more than one element at a time*/
	numbersSlice = append(numbersSlice, 2, 3, 4)
	printSlice(numbersSlice)

	numbersSlice = append(numbersSlice, numbersSlice...) // variable unboxing
	printSlice(numbersSlice)

	/* create a slice numbers1 with double the capacity of earlier slice*/
	numbers1 = make([]int, len(numbers), (cap(numbers))*2)

	/* copy content of numbers to numbers1 */
	copy(numbers1, numbersSlice)
	printSlice(numbers1)

	// Map
	var countryCapitalMap map[string]string
	/* create a map*/
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* test if entry is present in the map or not*/
	capital, ok := countryCapitalMap["United States"]

	/* if ok is true, entry is present otherwise entry is absent*/
	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	/* We can also write above code as follow */
	//if capital, ok := countryCapitalMap["United States"]; ok {
	//	/* if ok is true, entry is present otherwise entry is absent*/
	//	fmt.Println("Capital of United States is", capital)
	//} else {
	//	fmt.Println("Capital of United States is not present")
	//}
}

func swap(x, y string) (string, string) {
	return y, x
}

func swap_p(x, y *int) {
	var temp int
	temp = *x /* save the value at address x */
	*x = *y   /* put y into x */
	*y = temp /* put temp into y */
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
