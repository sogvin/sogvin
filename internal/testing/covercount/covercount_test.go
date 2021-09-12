package covercount

import (
	"io/ioutil"
	"testing"
)

func Test_sing(t *testing.T) {
	songs := []int{0, 1, 2, 3}
	for _, song := range songs {
		t.Run("", func(t *testing.T) {
			defer mustIncreaseCoverage()(t)
			sing(ioutil.Discard, song)
		})
	}
}

func mustIncreaseCoverage() func(t *testing.T) {
	coverage := testing.Coverage()
	return func(t *testing.T) {
		if coverage == testing.Coverage() {
			t.Helper()
			t.Fatal("no increase in coverage")
		}
	}
}
