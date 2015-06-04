package main

import (
	// include standard testing-package
	"testing"
)

func TestDummy(t *testing.T) {
	d := Dummy{21}
	if d.Number != 21 {
		// report
		t.Errorf("Expected '%+v', actual '%+v'", 21, d.Number)
	}
}

// TODO: add unit test to provve serialisation
