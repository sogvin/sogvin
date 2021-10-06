package spec

import (
	"testing"
)

func Test_index(t *testing.T) {
	err := NewSpecification().SaveAs("docs/index.html")
	if err != nil {
		t.Fatal(err)
	}
}
