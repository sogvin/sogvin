package notes

import (
	"testing"
)

func Test_generate_www(t *testing.T) {
	book := NewBook()
	book.SaveTo("./se")
}
