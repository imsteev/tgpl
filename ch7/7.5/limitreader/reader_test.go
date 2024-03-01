package limitreader

import (
	"fmt"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	sr := strings.NewReader("stops herethiswillnotberead")
	lr := LimitReader(sr, 10)
	b := make([]byte, 50)
	lr.Read(b)
	fmt.Println(b)
}
