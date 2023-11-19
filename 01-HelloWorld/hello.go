package main

import "fmt"

// This is a commentary
/* THis is a multi-line
commentary */

// Declare a variable use this sintaxe, "var <name> <type>", you can also init tha variable value in the same sentence
// bellow some global variables was declared.
var b int = 22
var c, d string = "Hello", "world"

func main() {
	// I can also declare variable only assing a valuem and the compiler/interpreter will resolver what type the variable has
	a := 10
	e := "Hello"
	f := 10.33
	g := false
	h := 'W'
	i := `test
		multi line
		string`

	// Print variable value using fmt.Printf
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)
	fmt.Printf("%v \n", i)

	// Print variable types using fmt.Printf
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
	fmt.Printf("%T \n", h)
	fmt.Printf("%T \n", i)
}
