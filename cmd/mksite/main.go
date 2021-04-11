package main

import (
	"flag"
	"log"

	"github.com/gregoryv/sogvin"
)

func main() {
	prefix := flag.String("p", "./docs", "write pages to")
	flag.Parse()

	book := sogvin.NewSoftwareEngineeringBook()
	log.SetFlags(0)
	err := book.SaveTo(*prefix)
	if err != nil {
		log.Fatal(err)
	}
}
