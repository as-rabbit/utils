package slicex

import (
	"testing"
)

func TestStringEmpty(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{[]string{"", "1", ""}, []string{"1"}},
		{[]string{"", "", ""}, []string{}},
		{[]string{}, []string{}},
	}
	
	for _, tc := range testCases {
		actual := SliceFilter(tc.input, StringEmpty[string])
		
		if len(actual) != len(tc.expected) {
			t.Errorf("input: %s, expected: %s, got: %s", tc.input, tc.expected, actual)
		}
	}
}

func TestIntEmpty(t *testing.T) {
	testCases := []struct {
		input    []int64
		expected []int64
	}{
		{[]int64{0, 1, 0}, []int64{1}},
		{[]int64{0, 0, 0}, []int64{}},
		{[]int64{}, []int64{}},
	}
	
	for _, tc := range testCases {
		actual := SliceFilter(tc.input, IntegerEmpty[int64])
		
		if len(actual) != len(tc.expected) {
			t.Errorf("input: %d, expected: %d, got: %d", tc.input, tc.expected, actual)
		}
	}
}

func TestFloatEmpty(t *testing.T) {
	testCases := []struct {
		input    []float64
		expected []float64
	}{
		{[]float64{0, 1, 0}, []float64{1}},
		{[]float64{0, 0, 0}, []float64{}},
		{[]float64{}, []float64{}},
	}
	
	for _, tc := range testCases {
		actual := SliceFilter(tc.input, IntegerEmpty[float64])
		
		if len(actual) != len(tc.expected) {
			t.Errorf("input: %f, expected: %f, got: %f", tc.input, tc.expected, actual)
		}
	}
}
