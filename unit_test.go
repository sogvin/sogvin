package sogvin

import (
	"testing"
)

func Test_generate_www(t *testing.T) {
	book := NewBook()
	err := book.SaveTo("./htdocs")
	if err != nil {
		t.Error(err)
	}
}
