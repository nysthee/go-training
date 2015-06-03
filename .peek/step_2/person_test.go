package main

import (
	// include standard testing-package
	"bytes"
	"reflect"
	"testing"
)

func TestPerson(t *testing.T) {
	p := Person{Age: 21}
	if p.Age != 21 {
		// report
		t.Errorf("Expected '%+v', actual '%+v'", 21, p.Age)
	}
}

func TestJson(t *testing.T) {
	p := Person{Name: "Marc", Age: 42, Interests: []string{"Running", "Cycling", "Hockey"}}

	var buffer bytes.Buffer
	err := toJson(p, &buffer)
	if err != nil {
		t.Errorf("Expected json encoding to succeed: %+v", err)
	}

	var again Person
	err = fromJson(&buffer, &again)
	if err != nil {
		t.Errorf("Expected json decoding to succeed: %+v", err)
	}

	equal := reflect.DeepEqual(p, again)
	if equal != true {
		t.Errorf("Expected '%+v' not equal to actual: '%+v", p, again)
	}
}
