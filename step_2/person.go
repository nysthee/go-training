package main

import
// import standard libraries or 3rd party packages

"log"

// TODO:
// rename this struct and
// add the following attributes:
// - your name
// - age and
// - interests (asslice)
//
type Dummy struct {
	Number int
}

func main() {
	// create and populate the struct
	d := Dummy{Number: 42}

	// log the struct and its attributes in internal verbose format
	log.Printf("%+v\n", d)

	// TODO print the json to stdout

}
