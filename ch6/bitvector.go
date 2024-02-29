package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	// [0-63], [64-127], [128-191], ... [i*64-(i*64+63)],
	words []uint64
}

// how to handle negatives?
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, x%64 // divide by 64 to get offset by multiple, and then modulo to get remainder bit
	return word < len(s.words) && (s.words[word]&(1<<bit)) > 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, x%64
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bit)
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tw := range t.words {
		if i < len(s.words) {
			s.words[i] |= tw

		} else {
			s.words = append(s.words, tw)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	var c int
	for _, word := range s.words {
		n := word
		for n > 0 {
			c++
			n &= (n - 1)
		}
	}
	return c
}

func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		return
	}
	word, bit := x/64, x%64
	s.words[word] &= (^(1 << bit)) // need to flip all these bits
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	var t IntSet
	t.UnionWith(s)
	return &t
}

func main() {
	var x, y IntSet
	x.AddAll(1, 144, 9)
	fmt.Println(x.String())
	y.AddAll(9, 42)
	fmt.Println(y.String())
	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123), x.Has(42))
	t := x.Copy()
	t.Add(555)
	fmt.Println(t.String())
	t.Remove(9)
	fmt.Println(x.Has(555), t.Has(555), x.Has(9), t.Has(9))
}
