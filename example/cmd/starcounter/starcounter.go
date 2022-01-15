package main

import (
	"fmt"
)

// NewStarCounter returns a new star counter command with logging
// enabled to cmd.Stderr.
func NewStarCounter() *StarCounter {
	return &StarCounter{}
}

type StarCounter struct {
	size   string
	weight int
}

func (s *StarCounter) SetSize(v string) { s.size = v }
func (s *StarCounter) SetWeight(v int)  { s.weight = v }

// Run starts the application and waits for it to complete.
func (s *StarCounter) Run() error {
	// count stars using filters
	if s.weight < 0 {
		return fmt.Errorf("bad weight")
	}
	// do the actual counting...
	return nil
}
