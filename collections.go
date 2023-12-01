package main

// Returns the index of e in s. -1 if s does not contain e
func IndexOf[T comparable](s []T, e T) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

// Returns true if s contains e
func Contains[T comparable](s []T, e T) bool {
	return IndexOf(s, e) != -1
}
