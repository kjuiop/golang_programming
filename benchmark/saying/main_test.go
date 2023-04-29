package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Jake")
	if s != "Welcome my deer Jake" {
		t.Error("got", s, "expected", "Welcome my deer Jake")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Jake"))
	// Output:
	// Welcome my dear Jake
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Jake")
	}
}
