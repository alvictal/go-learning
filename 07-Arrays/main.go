package main

import "fmt"

// Arrays works prettu straight forward as other languages, the exemple below let make things clear, but in go, arrays will always have
// a fixed size. For dynamic size check out the next lesson about slices
func printArray(arr []int, size int) {
	var i int
	for i = 0; i < size; i++ {
		fmt.Printf("i: %v , v: %v\n", i, arr[i])
	}
}

func main() {
	var x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x[0] = 1
	x[1] = 38
	x[9] = 39
	printArray(x, 10)

	y := []int{3, 4, 5, 6, 7}
	printArray(y, 5)

}
