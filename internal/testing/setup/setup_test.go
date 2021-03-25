package setup

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// setup before

	code := m.Run()
	defer os.Exit(code)

	// teardown after
}
