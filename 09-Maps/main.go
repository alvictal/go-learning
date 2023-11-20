package main

import "fmt"

// Maps are similar to dictionaries in other languages, we will use a combination of key-value pair to access the elements inside a map
func main() {
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 20

	fmt.Println(m)
	fmt.Println(m["a"])

	delete(m, "b")
	fmt.Println(m)

	m["c"] = 30
	fmt.Println(m)

	// below,  is the way that we check if a value exist inside a map
	_, exists := m["b"]
	fmt.Println(exists)

	// and for last, there is a another way to declare a map
	x := map[string]int{"a": 100, "b": 200}
	fmt.Println(x)

	//to finish lets clear the last map that we created
	clear(x)
	fmt.Println(x)
}
