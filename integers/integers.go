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
