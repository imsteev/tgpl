package treesort

import (
	"fmt"
	"testing"
)

func TestTreeString(t *testing.T) {
	tr := &tree{value: 5}
	add(tr, 1)
	add(tr, 2)
	add(tr, 9)
	add(tr, -1)
	add(tr, 0)
	add(tr, -500)
	fmt.Println(tr.String())
}
