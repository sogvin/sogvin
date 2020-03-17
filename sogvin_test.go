package sogvin

import "testing"

func Test_generate(t *testing.T) {
	book := NewSoftwareEngineeringBook()
	if err := book.SaveTo("./htdocs"); err != nil {
		t.Error(err)
	}
}
