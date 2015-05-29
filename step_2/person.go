package main

import (
	// import standard libraries or 3rd party packages
	"fmt"
)

// define your struct
type Dummy struct {
	number int
}

func main() {
	// create and populate the struct
	d := Dummy{number: 42}

	// print the struct and its attributes to stdout
	fmt.Printf("%+v\n", d)
}
