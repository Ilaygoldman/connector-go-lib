package connector

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	expected := a + b

	if got := Add2(a, b); got != expected {
		t.Errorf("Add(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}

func TestSubtract(t *testing.T) {
	a := 1
	b := 2
	expected := a - b

	if got := Subtract2(a, b); got != expected {
		t.Errorf("Subtract(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}