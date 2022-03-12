package main

import (
	"playground/examples"
	"testing"
)

func TestSimpleFunction(t *testing.T) {
	expected := "D"
	result := examples.Run()

	if expected[0] != result[0] {
		t.Error("The result does not match what we expected")
	}
}
