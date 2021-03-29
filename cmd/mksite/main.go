package main

import (
	"log"

	"github.com/gregoryv/sogvin"
)

func main() {
	book := sogvin.NewSoftwareEngineeringBook()
	log.SetFlags(0)
	err := book.SaveTo("./docs")
	if err != nil {
		log.Fatal(err)
	}
}
