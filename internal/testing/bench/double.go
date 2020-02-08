package bench

import (
	"fmt"
)

func double1(i int) (int, error) {
	if i < 0 {
		return 0, fmt.Errorf("double: i must be positive")
	}
	n := i * 2
	if n < i {
		return MAX, nil
	}
	return n, nil
}

func double2(result *int, i int) error {
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

func double3(result *int, i int) error {
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
