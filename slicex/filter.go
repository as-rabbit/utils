package slicex

// SliceFilter is a generic function that filters elements of a slice based on a provided filtering function.
func SliceFilter[T any](s []T, fn func(i int, v T) bool) []T {
	
	var (
		filtered []T
		length   = len(s)
	)
	
	for i := 0; i < length; i++ {
		
		if fn(i, s[i]) {
			
			filtered = append(filtered, s[i])
			
		}
		
	}
	
	return filtered
}
