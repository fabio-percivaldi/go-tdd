package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	want := 4

	if sum != want {
		t.Errorf("got '%d' want '%d'", sum, want)
	}
}

func assertAdditionIsCorrect(t *testing.T, got, want int, numbers [5]int) {
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
func TestMultipleAdder(t *testing.T) {

	t.Run("multiple adder return 15 with [1, 2, 3, 4, 5]", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := MultipleAdd(numbers)
		want := 15

		assertAdditionIsCorrect(t, got, want, numbers)
	})

	t.Run("multiple adder return 10 with [2, 2, 2, 2, 2]", func(t *testing.T) {
		numbers := [5]int{2, 2, 2, 2, 2}
		got := MultipleAdd(numbers)
		want := 10

		assertAdditionIsCorrect(t, got, want, numbers)
	})
}
