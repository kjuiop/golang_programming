# Go 언어의 유닛 테스트

---

- Go 언어에서는 언어 자체적으로 유닛 테스트나 벤치 마크 테스트를 지원합니다.
- 테스트 실행 구문
    - `go test .`
- 벤치마크 테스트 실행 구문
    - `go test -bench=. -benchmem`

### SampleCode

```go
func main() {
	fmt.Println("result : ", Concat("a", "bcd"))
}

func Concat(a, b string) string {
	return a + b
}
```

### SampleTest

```go

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
```

### Result

```go
goos: darwin
goarch: amd64
pkg: golang_programming/benchmark/concat
cpu: VirtualApple @ 2.50GHz
BenchmarkConcat-8        1000000                 49910 ns/op            503993 B/op          1 allocs/op
BenchmarkBuffer-8       179282912                6.019 ns/op            3 B/op               0 allocs/op
BenchmarkCopy-8         369248331                3.284 ns/op            1 B/op               0 allocs/op
BenchmarkEqualFold-8    26838782                 50.44 ns/op            0 B/op               0 allocs/op
```