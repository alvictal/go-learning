package main

import "fmt"

func sum2(a *int) int {
	*a = *a + 2
	return *a
}

/*
	Pointers works similar to C or C++, we can assing the memmory adress to another variable, access the value using "*"

passing as reference on functions, etc
*/
func main() {
	x := 10
	fmt.Printf("X initial value and address\n")
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("\nY value and address\n")
	y := &x
	fmt.Println(y)
	fmt.Println(*y)

	fmt.Printf("\nModifying a variable that habe the reference to another variable\n")
	*y = 20
	fmt.Println(x)

	fmt.Printf("\nDeclaring a pointer with defined type\n")
	var z *int = &x
	fmt.Println(z)
	fmt.Println(*z)

	fmt.Printf("\nUsing functions to pass variable as reference\n")
	b := 100
	sum2(&b)
	fmt.Println(b)
}
