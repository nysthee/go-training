package main

import (
	// import standard libraries or 3rd party packages
	"fmt"
)

// TODO: create a struct with your name, age and interests (where interests are moeled a string slice)
type Dummy struct {
	number int
}

func main() {
	// create and populate the struct
	d := Dummy{number: 42}

	// print the struct and its attributes to stdout
	fmt.Printf("%+v\n", d)
}
