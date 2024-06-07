package main

// Helper functions that in later version should use the package "slices". Currently using go versio 1.20

// func Index(s []string, v string) int {
// 	for i := range s {
// 		if v == s[i] {
// 			return i
// 		}
// 	}
// 	return -1

// }

func Index[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func Contains[S ~[]E, E comparable](s S, v E) bool {
	return Index(s, v) >= 0
}
