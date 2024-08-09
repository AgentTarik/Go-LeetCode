package problems

func Reverse(x int) int {

	reversed := 0

	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	if reversed < -2147483648 || reversed > 2147483647 {
		return 0
	}

	return reversed
}
