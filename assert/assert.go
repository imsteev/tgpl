package assert

import "log"

func Equal(expected, actual interface{}, msg ...string) {
	if expected != actual {
		log.Fatalf("expected %v, got %v", expected, actual)
	}
}
