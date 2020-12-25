package sogvin

import "testing"

func Test_generate(t *testing.T) {
	book := NewSoftwareEngineeringBook()
	if err := book.SaveTo("./docs"); err != nil {
		t.Error(err)
	}
}
