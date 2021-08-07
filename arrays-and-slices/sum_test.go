package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertEqual := func(t *testing.T, expected, got int, numbers []int) {
		t.Helper()
		if expected != got {
			t.Errorf("expected %d, received %d, given %v", expected, got, numbers)
		}
	}
	t.Run("collection of any length", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		got := Sum(numbers)
		expected := 28
		assertEqual(t, expected, got, numbers)
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, expected, got []int) {
		t.Helper()
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v, received %v", expected, got)
		}
	}
	t.Run("Sum all tails of given slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5})
		expected := []int{5, 5}
		checkSums(t, expected, got)
	})
	t.Run("gracefully handle empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 5})
		expected := []int{0, 5}
		checkSums(t, expected, got)
	})

}
