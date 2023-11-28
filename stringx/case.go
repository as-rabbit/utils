package stringx

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
)

// Upper converts the first letter of the input string to uppercase.
// If the input string is empty, an empty string is returned.
func Upper(s string) string {
	
	if "" == s {
		return ""
	}
	
	return cases.Title(language.English, cases.NoLower).String(s)
}

// FirstLower converts the first letter of the input string to lowercase
// If the input string is empty,an empty string is returned.
func FirstLower(s string) string {
	
	if len(s) == 0 {
		return s
	}
	
	return strings.ToLower(string(s[0])) + s[1:]
}

// PascalCase converts an underscore separated string
func PascalCase(s string) string {
	
	words := strings.Split(s, "_")
	
	if 1 == len(words) {
		return Upper(s)
	}
	
	for i, word := range words {
		if len(word) > 0 {
			words[i] = Upper(word)
		}
	}
	
	return strings.Join(words, "")
}

// CamelCase converts a string with underscores
func CamelCase(s string) string {
	
	words := strings.Split(s, "_")
	
	if 1 == len(words) {
		return FirstLower(s)
	}
	
	for i, item := range words {
		
		if 0 == i {
			words[i] = FirstLower(item)
			continue
		}
		
		words[i] = Upper(item)
	}
	return strings.Join(words, "")
}

// Snake converts a camelCase string to snake_case format.
func Snake(s string) string {
	
	var builder strings.Builder
	
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			builder.WriteRune('_')
		}
		builder.WriteRune(unicode.ToLower(r))
	}
	
	return builder.String()
}
