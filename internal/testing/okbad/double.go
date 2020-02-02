package okbad

import (
	"fmt"
)

// double returns the double of i if i is positive but never more than
// max int
func double(i int) (int, error) {
	if i < 0 {
		return 0, fmt.Errorf("double: i must be positive")
	}
	n := i * 2
	if n < i {
		return MAX, nil
	}
	return n, nil
}

const MAX int = int(^uint(0) >> 1)
