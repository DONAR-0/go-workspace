package arrayslices

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	// lengthOfNumbers := len(numbersToSum)
	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	// lengthOfNumbers := len(numbersToSum)
	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		numbers := numbers[1:]
		sums = append(sums, Sum(numbers))
	}

	return sums
}
