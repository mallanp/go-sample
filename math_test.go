package main

import (
	"testing"
)

func Test_add(t *testing.T) {

	actual := add(10, 20)
	if actual != 30 {
		t.Errorf("Expetced 30 but got %d", actual)
	}
}

func Test_prod(t *testing.T) {

	actual := prod(10, 20)
	if actual != 200 {
		t.Errorf("Expetced 200 but got %d", actual)
	}
}

func Test_fact(t *testing.T) {

	actual := fact(4)
	if actual != 24 {
		t.Errorf("Expetced 24 but got %d", actual)
	}
}
