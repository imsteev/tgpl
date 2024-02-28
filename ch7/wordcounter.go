package ch7

import (
	"bufio"
	"strings"
)

type WordCounter int

func (wc *WordCounter) Write(b []byte) (int, error) {
	var ct int
	r := strings.NewReader(string(b))
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		ct++
	}
	*wc = WordCounter(ct)
	return ct, nil
}