package main

// rotates clockwise
// "mangos" -> "angosm" (m[angosm]angos)
func rotate(s []int, n int) []int {
	n %= len(s)
	// the trick: put two slices next to each other and slide by # of rotations
	// return append(s, s...)[n : n+len(s)]

	// option 2: slice surgery
	chunk := s[:n]
	s = s[n:]
	s = append(s, chunk...)
	return s
}
