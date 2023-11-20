package main

import (
	"fmt"
	"strconv"
)

/*
	Functions works similar to other languages always with the structure "func <func_name>(<parameter>...) <return type> {}"". Void
	functions don't need to be typed (not necessary use void as return type). Typed functions will need return the a value related to the type, but depends
	how the fuction is construct there is no need to use the reserved keyword "return"
*/

func justReturn(a string) (x string) {
	x = a
	return
}

// We can return more the one variable in the function
func doNothing(a string, b int) (string, string) {

	return a, strconv.Itoa(b)
}

// We also can have indefined number of parameter with go
func variadicFunc(x ...int) int {
	var res int

	// This is an example of loop, in this case range return the index and the value
	// because we dont care about the index, we put "_" to just ignore the value
	for _, v := range x {
		res += v
	}

	return res
}

// we also can creat anonymus functions as in other languages (crap javascript)
// there are two ways to do on is using a function inside a function as bellow, the other way
// we can find on main function: atributing the func keyword to a variable
func anonymusFuncInsideFunc() func() int {
	x := 10
	return func() int {
		return x * x
	}
}

func main() {
	fmt.Println(justReturn("Hello World"))

	x, y := doNothing("Hi", 30)
	fmt.Printf("%v %v\n", x, y)

	fmt.Printf("variadic Function: %v\n", variadicFunc(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	//here is the other way to implement an anonymus functions
	z := 2
	a := func() {
		z += 2
	}

	b := anonymusFuncInsideFunc()
	a()

	fmt.Printf("%v %v", b(), z)
}
