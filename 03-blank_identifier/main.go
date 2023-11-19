package main

import (
	"fmt"
	"net/http"
)

func blank_id() {
	res, _ := http.Get("http://www.uol.com.br")

	fmt.Printf("%v\n", res)
}

func main() {
	// Go doesn't allow us to create a variable and not use it (compiler error), but in some moments we want to ignore a value return
	// for this we should use the blank identifier "_"
	res, err := http.Get("http://fullcycle.com.br") //http get will return two values

	if err != nil {
		fmt.Printf("%v err accessing the website\n", err)
	} else {
		fmt.Printf("%v\n", res)
	}

	blank_id()
}
