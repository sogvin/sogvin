package spec

import (
	"testing"
)

func Test_specification(t *testing.T) {
	NewSpecification().SaveAs("spec.html")
}
