package arrayslices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Test all the values after a[0]
func TestSumTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func BenchmarkSum(b *testing.B) {
	for b.Loop() {
		Sum([]int{b.N, b.N + 1, b.N - 2})
	}
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{b.N, b.N + 1, b.N - 2}, []int{b.N, b.N - 1, b.N + 2})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for b.Loop() {
		SumAllTails([]int{b.N, b.N + 1, b.N - 2}, []int{b.N, b.N - 1, b.N + 2})
	}
}

func ExampleSum() {
	sum := Sum([]int{5, 8})
	fmt.Print(sum)
	// Output: 13
}

func ExampleSumAll() {
	sumAll := SumAll([]int{5, 8}, []int{15, 13})
	fmt.Print(sumAll)
	// Output: [13 28]
}

func ExampleSumAllTails() {
	sumAllTails := SumAllTails([]int{5, 8, 1}, []int{15, 13, 2})
	fmt.Print(sumAllTails)
	// Output: [9 15]
}
