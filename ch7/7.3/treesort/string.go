package treesort

import (
	"fmt"
	"strings"
)

// 7.3
func (t *tree) String() string {
	if t == nil {
		return ""
	}
	s := fmt.Sprintf("%s %d %s", t.left.String(), t.value, t.right.String())
	return strings.TrimSpace(s)
}
