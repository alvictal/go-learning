package main

import "fmt"

// Constans using the same principle of mutable variables, the diference is the use of "=" for constants and ":=" during value assign
const xyz int = 22
const x = 10

const (
	aa string = "teste"
	bb        = 60
	cc int    = 567
)

func main() {
	fmt.Printf("%v %v %v %v %v\n", xyz, x, aa, bb, cc)
}
