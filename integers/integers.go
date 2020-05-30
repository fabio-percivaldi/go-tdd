package integers

// Add returns the sum of two integers
func Add(first, second int) int {
	return first + second
}

// MultipleAdd acceptc an array of int and return the sum
func MultipleAdd(numbers []int) (sum int) {
	for _, addend := range numbers {
		sum += addend
	}
	return
}

// SumAll accepts a variable number of slices and return a slice with the sum of every input slice
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, slice := range numbersToSum {
		sums = append(sums, MultipleAdd(slice))
	}
	return
}

// SumAllTails accepts a variable number of slices and return a slice with the sum of every input slice except the first number of the slice
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, slice := range numbersToSum {
		sums = append(sums, MultipleAdd(slice[1:]))
	}
	return
}
