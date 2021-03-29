package main

import (
	"strings"
	"testing"

	"github.com/gregoryv/wolf"
)

func TestStarCounter(t *testing.T) {
	exp := func(exitCode int, args ...string) {
		tc := strings.Join(args, " ")
		t.Run(tc, func(t *testing.T) {
			cmd := wolf.NewTCmd(args...)
			defer cmd.Cleanup()
			sc := NewStarCounter(cmd)
			got := sc.Run()
			if got != exitCode {
				t.Error(args, "got", got, "expected", exitCode)
			}
		})
	}

	exp(0, "name")
	exp(0, "name", "-h")
	exp(0, "name", "--help")
	exp(0, "name", "-size", "small")
	exp(1, "name", "-weight", "heavy")
	exp(1, "", "-nosuch")
}
