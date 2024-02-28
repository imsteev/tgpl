package main

// return true if s1 and s2 are anagrams of each other
// a string is an anagram of another string if it can be re-arranged to equal the other.
func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make(map[rune]int)

	for _, r := range s1 {
		counts[r]++
	}

	for _, r := range s2 {
		counts[r]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}
