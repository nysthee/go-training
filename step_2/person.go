package main

import (
	// import standard libraries or 3rd party packages
	"fmt"
	"log"
	"os"
)

// TODO: rename this struct and add the following attributes:
// - your name
// age and
// interests (asslice)
//
type Dummy struct {
	number int
}

func main() {
	// create and populate the struct
	d := Dummy{number: 42}

	// log the struct and its attributes
	log.Printf("%+v\n", d)

	// print the json to stdout
	fmt.Fprintf(os.Stdout, "%+v", toJson(d))
	fmt.Fprintf(os.Stdout, "%s", string(toJson(d)))

}

func toJson(d Dummy) []byte {
	// TODO convert to json
	return nil
}
