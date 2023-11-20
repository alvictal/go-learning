package main

import "fmt"

// Slice is similar to lists in python or another coomparasion arrays with dynamic allocation
func main() {
	//We can use the function make to create a slice with a pre-defined number of positions (all position will init with 0 )
	slice := make([]int, 5)
	slice[0] = 10

	fmt.Println(slice[0])
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(slice)

	// Another way to declare a slice
	sliceString := []string{
		"Hello",
		"World",
	}

	fmt.Println(sliceString[0])

	//THe index control is similar to python too
	slice2String := []string{
		"This",
		"is a",
		"story of",
		"a girl",
	}

	fmt.Println(slice2String[:2])  // "This is a"
	fmt.Println(slice2String[1:2]) // "is a"
	fmt.Println(slice2String[2:4]) // "story of a girl"
	fmt.Println(slice2String[2:])  // "story of a girl"
}
