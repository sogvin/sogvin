package main

import (
	"fmt"

	"github.com/gregoryv/notes/page"
)

//go:generate go run .
func main() {
	fmt.Println("Dog")
	page.WriteAllPages("./www")
}
