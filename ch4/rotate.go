package main

// rotates clockwise
// "mangos" -> "angosm" (m[angosm]angos)
func rotate(s []int, n int) []int {
	// the trick: put two slices next to each other and slide by # of rotations
	return append(s, s...)[n : n+len(s)]
}
