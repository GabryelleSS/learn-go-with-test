package main

import (
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

	t.Run("sum", func(t *testing.T) {
		numbersA := []int{0, 9}
		numbersB := []int{5, 13, 27}

		got := SumAll(numbersA, numbersB)
		want := []int{9, 45}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}