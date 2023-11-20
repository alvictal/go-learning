package main

import "fmt"

/*
In Go, interface work a bit different from other languages,
First isn't necessary to use any reserved key as "implements"
second if the interface and struct has the same signature, so we say that interface implements that struct
*/
type vehicle interface {
	start() string
}

type car struct {
	Name string
}

func (c car) start() string {
	return "The " + c.Name + " started the engine"
}

func startingTheEngine(v vehicle) string {
	return v.start()
}

func main() {
	c := car{"Alpha Tauri"}
	fmt.Println(startingTheEngine(c))
}
