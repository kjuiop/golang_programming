package main

import "testing"

func TestConcat(t *testing.T) {
	s := Concat("a1", "b2")
	if s != "a1b2" {
		t.Errorf("Concat(a1, b2) = %s; want a1b2", s)
	}
}
