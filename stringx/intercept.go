package stringx

// FirstStr returns the first character of the given string as a string.
// If the given string is empty, it returns an empty string.
func FirstStr(s string) string {
	
	if len(s) == 0 {
		return ""
	}
	
	return string(s[0])
}
