package stringx

import "testing"

func TestPascalCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"the_quick_brown_fox", "TheQuickBrownFox"},
		{"theQuickBrownFox", "TheQuickBrownFox"},
		{"TheQuickBrownFox", "TheQuickBrownFox"},
		{"", ""},
	}
	
	for _, tc := range testCases {
		actual := PascalCase(tc.input)
		if actual != tc.expected {
			t.Errorf("input: %s, expected: %s, got: %s", tc.input, tc.expected, actual)
		}
	}
}

func TestCamelCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"theQuickBrownFox", "theQuickBrownFox"},
		{"the_quick_brown_fox", "theQuickBrownFox"},
		{"TheQuickBrownFox", "theQuickBrownFox"},
		{"", ""},
	}
	
	for _, tc := range testCases {
		actual := CamelCase(tc.input)
		if actual != tc.expected {
			t.Errorf("input: %s, expected: %s, got: %s", tc.input, tc.expected, actual)
		}
	}
}

func TestSnake(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"theQuickBrownFox", "the_quick_brown_fox"},
		{"TheQuickBrownFox", "the_quick_brown_fox"},
		{"the_quick_brown_fox", "the_quick_brown_fox"},
		{"", ""},
	}
	
	for _, tc := range testCases {
		actual := Snake(tc.input)
		if actual != tc.expected {
			t.Errorf("input: %s, expected: %s, got: %s", tc.input, tc.expected, actual)
		}
	}
}
