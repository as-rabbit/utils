package stringx

import "testing"

func TestFirstStr(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		// 测试用例1：非空字符串
		{input: "Hello, World!", expected: "H"},
		// 测试用例2：空字符串
		{input: "", expected: ""},
		// 可以继续添加更多的测试用例...
	}
	
	for _, tc := range testCases {
		result := FirstStr(tc.input)
		if result != tc.expected {
			t.Errorf("FirstStr(%q) = %q, want %q", tc.input, result, tc.expected)
		}
	}
}
