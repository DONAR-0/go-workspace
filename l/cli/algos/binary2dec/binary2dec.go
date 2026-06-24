package binary2dec

func Convert(input int) int {
	// Get Second Number
	decimal := 0
	base := 1

	for input > 0 {
		lastDigit := input % 10
		decimal += lastDigit * base
		base *= 2
		input = input / 10
	}

	return decimal
}
