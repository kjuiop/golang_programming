package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestConcat(t *testing.T) {
	s := Concat("a1", "b2")
	if s != "a1b2" {
		t.Errorf("Concat(a1, b2) = %s; want a1b2", s)
	}
}

func BenchmarkConcat(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str += "x"
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
}

func BenchmarkCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0
	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], "x")
	}
}

func BenchmarkEqualFold(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strings.EqualFold("abc", "ABC")
		_ = strings.EqualFold("ABC", "ABC")
		_ = strings.EqualFold("laBcD", "lAbCd")
	}
}
