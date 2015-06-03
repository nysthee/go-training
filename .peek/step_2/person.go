package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Person struct {
	Name      string
	Age       int
	Interests []string
}

func main() {
	// create and populate the struct
	p := Person{Name: "Marc", Age: 42, Interests: []string{"Running", "Cycling", "Hockey"}}

	// log the struct and its attributes
	log.Printf("%+v\n", p)

	// print the json to stdout
	toJson(p, os.Stdout)

}

func toJson(p Person, w io.Writer) error {

	enc := json.NewEncoder(w)
	err := enc.Encode(p)
	if err != nil {
		return err
	}
	return nil
}

func fromJson(r io.Reader, person *Person) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(&person)
	if err != nil {
		return err
	}
	return nil
}
