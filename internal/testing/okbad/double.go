package okbad

import (
	"fmt"
)

// double sets the result to the double of i if i is positive but
// never more than max int
func double(result *int, i int) error {
	if result == nil {
		return fmt.Errorf("double: result cannot be nil")
	}
	if i < 0 {
		*result = 0
		return fmt.Errorf("double: i must be positive")
	}
	n := i * 2
	if n < i {
		*result = MAX
		return nil
	}
	*result = n
	return nil
}

const MAX int = int(^uint(0) >> 1)
