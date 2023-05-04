package example

// MySum 함수는 무제한의 인수를 더한다.
func MySum(xi ...int) int {
	sum := 0
	for _, val := range xi {
		sum += val
	}
	return sum
}
